package service

import (
	"dayz-tool/global"
	"dayz-tool/model/common"
)

type ConfigService struct {
}

func (c *ConfigService) GetLocalConfig() (config common.LocalConfig) {
	global.DT_DB.Model(&common.LocalConfig{}).First(&config)
	return config
}

func (c *ConfigService) InitLocalConfig() {
	config := &common.LocalConfig{
		Theme: "light",
	}
	global.DT_DB.Create(config)

}
