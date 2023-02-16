package handler

import "github.com/gin-gonic/gin"

type ContentHandler interface {
	Create(c *gin.Context)
	FindByUserIdAndContentIdDetail(c *gin.Context)
}
