package main

import (
	"ginlearning/controller"
	"ginlearning/database"
	"ginlearning/middleware"
	"ginlearning/repository"
	"ginlearning/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := database.NewDb()
	validator := validator.New()
	categoryRepository := repository.NewRepository()
	categoryService := service.NewServices(categoryRepository, db, validator)
	categoryController := controller.NewControlers(categoryService)
	middlewares := middleware.NewMiddleware()

	r := gin.Default()
	r.POST("/login", categoryController.Login)
	r.POST("/regis", categoryController.Register)
	r.GET("/ping", middlewares.Midware, Ping)
	r.Run()
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
