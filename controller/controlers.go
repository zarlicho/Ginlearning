package controller

import "github.com/gin-gonic/gin"

type Controllers interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
