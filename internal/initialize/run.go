package initialize

import (
	"fmt"

	"github.com/cuonggosu97/go-ecommerce-backen-api/global"
	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("oke", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8082")
}
