package main

import (
	"log"
	"net/http"
	"restAPI_lms/auth"
	courses "restAPI_lms/course"
	"restAPI_lms/handler"
	"restAPI_lms/helper"
	"restAPI_lms/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	courseRepository := courses.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	courseService := courses.NewService(courseRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email-check", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddlerware(authService, userService), userHandler.UploadAvatar)
	api.POST("/users-update", authMiddlerware(authService, userService), userHandler.UpdateUser)

	api.GET("/courses", courseHandler.GetCourses)

	router.Run()
}

func authMiddlerware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Monyet") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenSplit := strings.Split(authHeader, " ")
		if len(tokenSplit) == 2 {
			tokenString = tokenSplit[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
