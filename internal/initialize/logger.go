package initialize

import (
	"github.com/cuonggosu97/go-ecommerce-backen-api/global"
	"github.com/cuonggosu97/go-ecommerce-backen-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
