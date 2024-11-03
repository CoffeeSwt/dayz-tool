package initialize

import (
	"context"
	"dayz-tool/global"
	"dayz-tool/logger"
	"dayz-tool/model"
)

func Init(ctx *context.Context) {
	dBManager := NewGlobalDB(ctx)
	global.DT_DB = dBManager.db
	global.DT_Logger = logger.NewLogger(ctx)
	global.APP_Context = ctx
	model.MigrateTable(dBManager.db)
}
