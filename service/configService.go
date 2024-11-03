package service

import (
	"dayz-tool/global"
	"dayz-tool/model/common"
	"errors"
	"os/exec"

	"gorm.io/gorm"
)

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

type ConfigService struct {
}

func (c *ConfigService) GetServerStartConfig() (config *common.ServerStartConfig) {
	result := global.DT_DB.Model(&common.ServerStartConfig{}).First(&config)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return config
	}
	defaultConfig := &common.ServerStartConfig{
		DayzServerPath: "",
	}
	return defaultConfig
}

func (c *ConfigService) GetStartCommand() (cmd *exec.Cmd) {
	config := c.GetServerStartConfig()

	return exec.Command(config.DayzServerPath)
}

func (c *ConfigService) getPortParm(config common.ServerStartConfig) (portParm string) {
	return "-port=" + config.Port
}

func (c *ConfigService) getClientModsParm(config common.ServerStartConfig) (clientModsParm string) {
	base := "-mod="
	return base
}

func (c *ConfigService) getServerModsParm(config common.ServerStartConfig) (serverModsParm string) {
	base := "-servermod="
	return base
}

//插入一条数据
// func (c *ConfigService) InitServerStartConfig() {
// 	config := &common.ServerStartConfig{}
// 	global.DT_DB.Create(config)
// }
