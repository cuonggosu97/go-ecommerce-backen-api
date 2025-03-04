package global

import (
	"github.com/cuonggosu97/go-ecommerce-backen-api/pkg/logger"
	"github.com/cuonggosu97/go-ecommerce-backen-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
