package service

import (
	"context"
	"dayz-tool/global"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewPathChooser() *PathChooser {
	return &PathChooser{}
}

type PathChooser struct {
	ctx *context.Context
}

func (p *PathChooser) BindCtx(ctx *context.Context) {
	p.ctx = ctx
}

func (p *PathChooser) ChooseFile() string {
	filePath, err := runtime.OpenFileDialog(*p.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		global.DT_Logger.Info(err.Error())
		return ""
	}
	global.DT_Logger.Info(filePath)
	return filePath
}

func (p *PathChooser) ChooseFileFolder() string {
	dirPath, err := runtime.OpenDirectoryDialog(*p.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		global.DT_Logger.Info(err.Error())
		return ""
	}
	global.DT_Logger.Info(dirPath)
	return dirPath
}
