package handler

import (
	"net/http"
	"task2/model/api"
	"task2/model/request"
	"task2/service"

	"github.com/gin-gonic/gin"
)

type userContentHandler struct {
	userContentService service.UserContentService
}

func NewUserContentHandler(userContentService service.UserContentService) *userContentHandler {
	return &userContentHandler{
		userContentService: userContentService,
	}
}

func (handler *userContentHandler) PostReward(c *gin.Context) {
	var req request.RequestCreateUserContent

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userContent, err := handler.userContentService.PostReward(c, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   userContent,
	})
}
