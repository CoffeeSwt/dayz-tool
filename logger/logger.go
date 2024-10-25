package logger

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Logger struct {
	ctx context.Context
}

func NewLogger(ctx context.Context) *Logger {
	return &Logger{ctx: ctx}
}

// Debug implements logger.Logger.
func (l *Logger) Debug(message string) {
	runtime.LogDebug(l.ctx, message)
}

// Error implements logger.Logger.
func (l *Logger) Error(message string) {
	runtime.LogError(l.ctx, message)
}

// Fatal implements logger.Logger.
func (l *Logger) Fatal(message string) {
	runtime.LogFatal(l.ctx, message)
}

// Info implements logger.Logger.
func (l *Logger) Info(message string) {
	runtime.LogInfo(l.ctx, message)
}

// Print implements logger.Logger.
func (l *Logger) Print(message string) {
	runtime.LogPrint(l.ctx, message)
}

// Trace implements logger.Logger.
func (l *Logger) Trace(message string) {
	runtime.LogTrace(l.ctx, message)
}

// Warning implements logger.Logger.
func (l *Logger) Warning(message string) {
	runtime.LogWarning(l.ctx, message)
}
