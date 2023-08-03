package handler

import (
	"net/http"
	"restAPI_lms/helper"
	"restAPI_lms/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to register", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	registeredUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token := "blabla"
	formatter := user.FormatUser(registeredUser, token)

	response := helper.APIResponse("Register successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
