package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db) // buat koneksi database di main,parsing ke new repository
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Tes simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "anak band"
	userInput.Password = "password"

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1") // API versioning

	api.POST("/users", userHandler.RegisterUser)

	router.Run() // execute

	// userService.RegisterUser(userInput)
	// user := user.User{Name: "Test simpan"}
	// userRepository.Save(user)

	// fmt.Println("Connection to database is good")

	// var users []user.User // struct User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	// fmt.Println(user.Email)
	// 	fmt.Println("================================================================")
	// }

	// router
	// router := gin.Default()
	// router.GET("/handler", handler) // mengarah ke function handler
	// router.Run()                    //

}

// func handler(c *gin.Context) { // handler untuk gin,c merupakan variabel
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// var users []user.User
// 	// db.Find(&users) // load struct User

// 	// c.JSON(http.StatusOK, users) // JSON

// 	//handler
// 	//service
// 	//repository
// 	//db

// }
