package initialize

import (
	"fmt"
	"net/http"

	c "github.com/cuonggosu97/go-ecommerce-backen-api/internal/controller"
	"github.com/cuonggosu97/go-ecommerce-backen-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Before ->> AA \n")
		c.Next()
		fmt.Printf("After ->> AA \n")
	}
}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("Before ->> BB \n")
		c.Next()
		fmt.Printf("After ->> BB \n")
	}
}
func CC(c *gin.Context) {
	fmt.Printf("Before ->> CC \n")
	c.Next()
	fmt.Printf("After ->> CC \n")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	//use the middleware
	r.Use(middlewares.AuthMiddleware(), AA(), BB(), CC)

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
