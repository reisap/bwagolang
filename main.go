package main

import (
	"fmt"
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

	input := user.LoginInput{
		Email: "postman@mail.com",
		Password: "lucubgt",
	}

	user, err := userService.Login(input)

	if err != nil {
		fmt.Println("user tidak ditemukan")
		fmt.Println(err.Error())
	}
	fmt.Println(user.Email)
	fmt.Println(user.Name)

	//userByEmail, err := userRepository.FindByEmail("r.prasapatraay@gmail.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(userByEmail.Name)

	router := gin.Default()
	api := router.Group("api/v1")
	api.POST("/users",userHandler.RegisterUser)

	router.Run()


	//input
	//handler
	//services : melakukan mapping dari struct user ke struct repository
	//repository
	//db
}