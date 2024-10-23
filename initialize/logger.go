package initialize

import (
	"dayz-tool/global"

	"go.uber.org/zap"
)

func initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	global.DT_Logger = logger
}
