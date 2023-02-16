package handler

import "github.com/gin-gonic/gin"

type UserContentHandler interface {
	PostReward(c *gin.Context)
}
