package initialize

import (
	"context"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBManager struct {
	ctx *context.Context
	db  *gorm.DB
}

func NewGlobalDB(ctx *context.Context) *DBManager {
	dbpath := getDataSource()
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &DBManager{ctx: ctx, db: db}
}

// 获取db文件夹路径
func getDataSource() string {
	cacheDir := getAbsPath()
	dataDir := filepath.Join(cacheDir, "dayz-tool")
	if err := os.MkdirAll(dataDir, os.FileMode(0755)); err != nil {
		return "dayz-tool.db"
	}
	return filepath.Join(dataDir, "dayz-tool.db")
}

func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	// fmt.Println("absPath:", dir)
	return dir
}
