package initialize

import (
	"context"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBManager struct {
	ctx context.Context
	db  *gorm.DB
}

func NewGlobalDB(ctx context.Context) *DBManager {
	dbpath := getDataSource()
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DBManager{ctx: ctx, db: db}
}

func (d *DBManager) MigrateTable(tables ...any) error {
	return d.db.AutoMigrate(tables...)
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
