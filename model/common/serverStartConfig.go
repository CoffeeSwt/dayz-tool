package common

type ServerStartConfig struct {
	DayzServerPath string `json:"dayzserverpath"`
	ClientMods     string `json:"clientmods"`
	ServerMods     string `json:"servermods"`
	Map            string `json:"map"`
}
