package glog

import (
	"log"
	"log/slog"
	"os"
	"strings"
	"syscall"
)

const defaultLevel = "INFO"

const envName = "LOG_LEVEL"

var levelMap = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

// getLogLevel returns the log level from the string
// 从给定的字符串中获取日志级别
func getLogLevel(l string) slog.Level {
	return levelMap[strings.ToUpper(l)]
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
	log.Printf("Finished to setup slog, LOG_LEVEL is %s \n", strings.ToUpper(level))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
