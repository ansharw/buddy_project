package main

import (
	"task2/database"
	"task2/handler"
	"task2/repository"
	"task2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	db := database.GetConnection()
	validate := validator.New()

	repoUser := repository.NewUserRepository()
	repoUserContent := repository.NewUserContentRepository()
	repoContent := repository.NewContentRepository()

	serviceUserContent := service.NewUserContentService(db, repoContent, repoUserContent, repoUser, *validate)
	serviceUser := service.NewUserService(db, repoUser, *validate)
	serviceContent := service.NewContentService(db, repoContent, repoUserContent, repoUser, *validate)

	handlerUser := handler.NewUserHandler(serviceUser)
	handlerUserContent := handler.NewUserContentHandler(serviceUserContent)
	handlerContent := handler.NewContentHandler(serviceContent)

	api.POST("/register", handlerUser.Register)
	api.POST("/login", handlerUser.Login)

	api.GET("/user/:user_id/content/:content_id", handlerContent.FindByUserIdAndContentIdDetail)
	api.POST("/user/:user_id/content", handlerContent.Create)
	api.GET("/user/:user_id/point", handlerUser.FindPointByUserId)

	api.GET("/company/registered_user", handlerUser.FindAllRegisteredUser)
	api.GET("/company/point", handlerUser.FindAllUserPoint)
	api.POST("/company/reward", handlerUserContent.PostReward)

	r.Run("localhost:9000")
}
