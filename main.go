package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangbwa/auth"
	"golangbwa/campaign"
	"golangbwa/handler"
	"golangbwa/helper"
	"golangbwa/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	dsn := "anaklucu:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	//repository
	userRepository := user.NewRepository(db)
	campaignsRepository := campaign.NewRepository(db)
	//service
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	campaignService := campaign.NewService(campaignsRepository)

	fmt.Println(campaignService)


	//handler
	userHandler := handler.NewUserHandler(userService,authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)


	router := gin.Default()
	//untuk dapat akses gambar yang disimpan di server
	router.Static("/images","./images") // ./images -> ini folder langsung di server,klon manya lucu maka ./lucu

	api := router.Group("api/v1")
	api.POST("/users",userHandler.RegisterUser)
	api.POST("/sessions",userHandler.Login)
	api.POST("/email_checkers",userHandler.CheckEmailAvailablity)
	api.POST("/avatar",authMiddleware(authService,userService),userHandler.UploadAvatars)

	//campaign
	api.GET("/campaign",campaignHandler.GetCampaign)
	router.Run()


	//input
	//handler
	//services : melakukan mapping dari struct user ke struct repository
	//repository
	//db
}

func authMiddleware(authService auth.Service,userService user.Service) gin.HandlerFunc{
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader,"Bearer"){
			response := helper.APIResponse("Unauthorized",http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader," ")
		if len(arrToken) == 2{
			tokenString = arrToken[1]
		}
		token,err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized",http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		claim,ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized",http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user,err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized",http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		c.Set("currentUser",user)

	}
}