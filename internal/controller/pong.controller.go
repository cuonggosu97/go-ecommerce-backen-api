package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println(("-----> My handler"))
	name := c.DefaultQuery("name", "cuongpham")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong..hhh..ping" + name,
		"uid":     uid,
		"user":    []string{"cr7", "m10", "cuongpham"},
	})
}
