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
	dbpath := getDataSource()
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

// 获取db文件夹路径
func getDataSource() string {
	cacheDir, _ := os.UserCacheDir()
	dataDir := filepath.Join(cacheDir, "dayz-tool")
	if err := os.MkdirAll(dataDir, os.FileMode(0755)); err != nil {
		return "dayz-tool.db"
	}

	return filepath.Join(dataDir, "dayz-tool.db")
}
