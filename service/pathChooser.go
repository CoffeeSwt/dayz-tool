package service

import (
	"dayz-tool/global"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewPathChooser() *PathChooser {
	return &PathChooser{}
}

type PathChooser struct {
}

func (p *PathChooser) ChooseFile() string {
	filePath, err := runtime.OpenFileDialog(*global.APP_Context, runtime.OpenDialogOptions{})
	if err != nil {
		global.DT_Logger.Info(err.Error())
		return ""
	}
	global.DT_Logger.Info(filePath)
	return filePath
}

func (p *PathChooser) ChooseFileFolder() string {
	dirPath, err := runtime.OpenDirectoryDialog(*global.APP_Context, runtime.OpenDialogOptions{})
	if err != nil {
		global.DT_Logger.Info(err.Error())
		return ""
	}
	global.DT_Logger.Info(dirPath)
	return dirPath
}
