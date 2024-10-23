package initialize

import (
	"dayz-tool/global"
	"dayz-tool/model/common"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func initGormDataBaseConnection() {
	workingDir, e := os.Getwd()
	if e != nil {
		return
	}
	dbpath := filepath.Join(workingDir, "db", "dayztool.db")
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		global.DT_DB = db

	}

}

func registerTables() {
	db := global.DT_DB
	err := db.AutoMigrate(
		common.LocalConfig{},
	)
	if err != nil {
		global.DT_Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	//global.DT_Logger.Info("register table success")
}
