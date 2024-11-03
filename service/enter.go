package service

func InitService() []any {
	return []any{
		NewPathChooser(),
		NewServerManager(),
		NewConfigService(),
	}
}
