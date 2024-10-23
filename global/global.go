package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DT_DB     *gorm.DB
	DT_Logger *zap.Logger
)
