package handler

import (
	"net/http"
	"restAPI_lms/course"
	"restAPI_lms/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type courseHandler struct {
	service course.Service
}

func NewCourseHandler(service course.Service) *courseHandler {
	return &courseHandler{service}
}

func (h *courseHandler) GetCourses(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	division := c.Query("division")

	courses, err := h.service.GetCourses(userID, division)
	if err != nil {
		response := helper.APIResponse("Failed to get courses", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of courses", http.StatusOK, "success", courses)
	c.JSON(http.StatusOK, response)
}
