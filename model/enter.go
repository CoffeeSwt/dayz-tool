package model

import (
	"dayz-tool/global"
	"dayz-tool/model/common"

	"gorm.io/gorm"
)

func MigrateTable(db *gorm.DB) {
	err := db.AutoMigrate(
		//add tables model here
		&common.ServerStartConfig{},
	)
	if err != nil {
		global.DT_Logger.Error(err.Error())
	}
}
