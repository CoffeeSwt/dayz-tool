package global

import (
	"dayz-tool/logger"

	"gorm.io/gorm"
)

var (
	DT_DB     *gorm.DB
	DT_Logger *logger.Logger
)
