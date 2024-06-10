package routers

import (
	"net/http"

	c "github.com/cuonggosu97/go-ecommerce-backen-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserByID)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "cuongpham")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong..hhh..ping" + name,
		"uid":     uid,
		"user":    []string{"cr7", "m10", "cuongpham"},
	})
}
