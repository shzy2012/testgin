package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shzy2012/testgin/service/aotm"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (t *HomeController) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"who": "home",
	})
}

func (t *HomeController) SetNX(c *gin.Context) {

	aotmService := aotm.NewAotmService()
	aotmService.AotmService()

	c.JSON(200, gin.H{
		"who":     "SetNX",
		"message": "ok",
	})

}
