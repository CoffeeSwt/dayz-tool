package common

import "gorm.io/gorm"

type LocalConfig struct {
	gorm.Model
	Theme string
}
