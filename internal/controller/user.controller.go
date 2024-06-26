package controller

import (
	"github.com/cuonggosu97/go-ecommerce-backen-api/internal/service"
	"github.com/cuonggosu97/go-ecommerce-backen-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {

	// response.SuccessResponse(c, 20001, []string{"cr7", "m10", "cuongpham"})
	response.ErrorResponse(c, 20003, "loi roi")
}
