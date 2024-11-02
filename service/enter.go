package service

import (
	"context"
)

type Service interface {
	BindCtx(*context.Context)
}

func InitService() []Service {
	return []Service{
		NewPathChooser(),
		NewServerManager(),
		NewConfigService(),
	}
}
