package glog

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"syscall"
)

const defaultLevel = "info"

const envName = "LOG_LEVEL"

var levelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

// getLogLevel returns the log level from the string
// 从给定的字符串中获取日志级别
func getLogLevel(l string) slog.Level {
	return levelMap[strings.ToLower(l)]
}

// init import this package to initialize the logger and setup slog
func init() {
	level, found := syscall.Getenv(envName)
	if !found {
		level = defaultLevel
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true, Level: getLogLevel(level),
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func Debug(v any) {
	slog.Debug(fmt.Sprint(v))
}

func Info(v any) {
	slog.Info(fmt.Sprint(v))
}

func Warn(v any) {
	slog.Warn(fmt.Sprint(v))
}

func Error(v any) {
	slog.Error(fmt.Sprint(v))
}
