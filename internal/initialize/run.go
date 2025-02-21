package initialize

import (
	"fmt"

	"github.com/cuonggosu97/go-ecommerce-backen-api/global"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8082")
}
