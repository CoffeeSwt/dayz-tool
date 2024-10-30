package service

import (
	"dayz-tool/global"
	"dayz-tool/model/common"
)

type ConfigService struct {
}

func (c *ConfigService) GetServerStartConfig() (config common.ServerStartConfig) {
	global.DT_DB.Model(&common.ServerStartConfig{}).First(&config)
	return config
}

func (c *ConfigService) InitServerStartConfig() {
	config := &common.ServerStartConfig{}
	global.DT_DB.Create(config)
}
