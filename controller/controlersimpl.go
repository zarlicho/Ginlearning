package controller

import (
	"ginlearning/model/web"
	"ginlearning/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControlersImpl struct {
	Service service.Service
}

func NewControlers(Services service.Service) Controllers {
	return &ControlersImpl{Service: Services}
}

func (Controler *ControlersImpl) Login(c *gin.Context) {
	var body web.LoginRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})
	}
	token, err := Controler.Service.Login(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to login!",
		})
		panic("failed to login")
	}
	//create cookies by token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token.Token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"info": "success to login!",
	})
}

func (Controler *ControlersImpl) Register(c *gin.Context) {
	var body web.RegisterRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})
	}

	Controler.Service.Register(c.Request.Context(), body)
	c.JSON(http.StatusOK, gin.H{
		"info": "success to regis!",
	})
}
