package handler

import (
	"net/http"
	"strconv"
	"task2/model/api"
	"task2/model/request"
	"task2/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (handler *userHandler) Login(c *gin.Context) {
	var req request.RequestLoginUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := handler.userService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Status: "success",
		Code:   http.StatusOK,
	})
}

func (handler *userHandler) Register(c *gin.Context) {
	var req request.RequestRegisterUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := handler.userService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, api.ApiResponse{
		Status: "success",
		Code:   http.StatusOK,
	})
}

func (handler *userHandler) FindAllRegisteredUser(c *gin.Context) {
	RegisteredUser, err := handler.userService.FindAllRegisteredUser(c)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   RegisteredUser,
	})
}

func (handler *userHandler) FindAllUserPoint(c *gin.Context) {
	AllUserPoint, err := handler.userService.FindAllUserPoint(c)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   AllUserPoint,
	})
}

func (handler *userHandler) FindPointByUserId(c *gin.Context) {
	val, ok := c.Params.Get("user_id")
	if !ok {
		panic("id not found")
	}
	user_id, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	UserWithPoint, err := handler.userService.FindPointByUserId(c, user_id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   UserWithPoint,
	})
}
