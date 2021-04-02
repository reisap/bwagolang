package main
import (
	"fmt"
	"golangbwa/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Reisa30041989@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection to database is good")
	var users []user.User
	db.Find(&users)

	length := len(users)
	fmt.Println(length)

	for _,user := range(users){
		fmt.Println(user.Email)
		fmt.Println(user.Name)
	}
}