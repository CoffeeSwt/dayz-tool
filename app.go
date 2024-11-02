package main

import (
	"context"
	"dayz-tool/initialize"
	"dayz-tool/service"
	"fmt"
)

// App struct
type App struct {
	ctx      context.Context
	services []service.Service
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		services: service.InitService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	initialize.Init(&ctx)
	for _, v := range a.services {
		v.BindCtx(&ctx)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
