package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "TipsGo", 40) // like fmt.Printf(format, args...)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40)) // like fmt.Println(args...)

	// 2
	// logger := zap.NewExample()
	// logger.Info("Hello")
	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello New Development")

	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello New Production")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// 1734342374.908031 -> 2024-12-16T16:46:14.907+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
