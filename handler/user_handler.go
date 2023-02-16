package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	FindAllRegisteredUser(c *gin.Context)
	FindAllUserPoint(c *gin.Context)
	FindPointByUserId(c *gin.Context)
}
