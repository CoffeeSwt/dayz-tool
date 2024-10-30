package service

import (
	"dayz-tool/global"
	"dayz-tool/model/common"
	"os/exec"
)

type ConfigService struct {
	config *common.ServerStartConfig
}

func (c *ConfigService) GetServerStartConfig() (config *common.ServerStartConfig) {
	global.DT_DB.Model(&common.ServerStartConfig{}).First(&config)
	c.config = config
	return config
}

func (c *ConfigService) GetStartCommand() (cmd *exec.Cmd) {
	config := c.GetServerStartConfig()

	return exec.Command(config.DayzServerPath)
}

func (c *ConfigService) getPortParm() (portParm string) {
	return "-port=" + c.config.Port
}

func (c *ConfigService) getClientModsParm() (clientModsParm string) {
	base := "-mod="
	return base
}

func (c *ConfigService) getServerModsParm() (serverModsParm string) {
	base := "-servermod="
	return base
}

//插入一条数据
// func (c *ConfigService) InitServerStartConfig() {
// 	config := &common.ServerStartConfig{}
// 	global.DT_DB.Create(config)
// }
