package main

import (
	"github.com/gin-gonic/gin"
	"golangbwa/handler"
	"golangbwa/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:Reisa30041989@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users",userHandler.RegisterUser)
	api.POST("/sessions",userHandler.Login)
	api.POST("/email_checkers",userHandler.CheckEmailAvailablity)
	api.POST("/avatar",userHandler.UploadAvatars)

	router.Run()


	//input
	//handler
	//services : melakukan mapping dari struct user ke struct repository
	//repository
	//db
}