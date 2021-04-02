package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangbwa/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//router := gin.Default()
	//router.GET("/",handler)
	//router.Run()
	dsn := "root:Reisa30041989@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.Name = "dari service"
	userInput.Password= "halo bro"
	userInput.Occupation ="Petani"
	userInput.Email = "petani@gmail.com"

	userService.RegisterUser(userInput)


	//contoh cek tanpa service
	//user := user.User{
	//	Name:"gorun",
	//}
	//userRepository.Save(user)

	//input
	//handler
	//services : melakukan mapping dari struct user ke struct repository
	//repository
	//db
}

func handler(c *gin.Context){
	dsn := "root:Reisa30041989@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection to database is good")
	var users []user.User
	db.Find(&users)

	for _,user := range(users){
		fmt.Println(user.Email)
		fmt.Println(user.Name)
	}

	c.JSON(http.StatusOK,users)

}