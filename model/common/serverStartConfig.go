package common

type ServerStartConfig struct {
	DayzServerPath string `json:"dayzserverpath"`
	ClientModsPath string `json:"clientmods"`
	ServerModsPath string `json:"servermods"`
	Port           string `json:"port"`
}
