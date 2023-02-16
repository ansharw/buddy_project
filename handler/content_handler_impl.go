package handler

import (
	"net/http"
	"strconv"
	"task2/model/api"
	"task2/model/request"
	"task2/service"

	"github.com/gin-gonic/gin"
)

type contentHandler struct {
	contentService service.ContentService
}

func NewContentHandler(contentService service.ContentService) *contentHandler {
	return &contentHandler{
		contentService: contentService,
	}
}

func (handler *contentHandler) Create(c *gin.Context) {
	val, ok := c.Params.Get("user_id")
	if !ok {
		panic("id not found")
	}
	user_id, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	var req request.RequestCreateContent

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ContentCreated, err := handler.contentService.Create(c, req, user_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   ContentCreated,
	})
}

func (handler *contentHandler) FindByUserIdAndContentIdDetail(c *gin.Context) {
	val, ok := c.Params.Get("user_id")
	if !ok {
		panic("user id not found")
	}

	user_id, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	val_, ok := c.Params.Get("content_id")
	if !ok {
		panic("content id not found")
	}

	content_id, err := strconv.Atoi(val_)
	if err != nil {
		panic(err)
	}
	content, err := handler.contentService.FindContentByUserId(c, user_id, content_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   content,
	})
}
