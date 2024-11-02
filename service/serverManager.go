package service

import (
	"context"
	"dayz-tool/global"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func NewServerManager() *ServerManager {
	return &ServerManager{}
}

type ServerManager struct {
	ctx    *context.Context
	cancel context.CancelFunc
}

func (s *ServerManager) BindCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *ServerManager) StartServer() {
	configService := &ConfigService{}
	cmd := configService.GetStartCommand()
	var wg *sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	wg.Add(1)
	//启动一个协程
	go s.runExe(cmd, ctx, wg)
}

func (s *ServerManager) CancelServer() {
	s.cancel()
}

func (s *ServerManager) runExe(cmd *exec.Cmd, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping the execution of the executable.")
			return
		default:
			// 启动命令
			if err := cmd.Start(); err != nil {
				global.DT_Logger.Error("Error starting command:" + err.Error())
				return
			}
			global.DT_Logger.Info("Command started, PID:" + string(cmd.Process.Pid))

			// 等待命令完成
			if err := cmd.Wait(); err != nil {
				global.DT_Logger.Error("Command crashed or ReStart with error:" + err.Error())
			}

			// 这里可以添加延迟，避免快速重启造成过大的负担
			global.DT_Logger.Info("Restarting Server...")
			time.Sleep(2 * time.Second) // 等待2秒后重启
		}
	}
}
