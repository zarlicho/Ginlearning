package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middlewares interface {
	Midware(c *gin.Context)
}
