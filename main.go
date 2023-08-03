package main

import (
	"log"
	"restAPI_lms/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_lms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	user := user.RegisterUserInput{
		Name:         "admin",
		Nim:          "21110395",
		Email:        "admin@gmail.com",
		Division:     "Pemograman",
		AlasanDaftar: "Mau Kaya",
	}

	userService.RegisterUser(user)
}
