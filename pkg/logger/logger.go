package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/easybyr/go-line/common/util"
)

var (
	_logger *slog.Logger
)

const (
	LOG_DIR  = "/log"
)

func init() {
	initLogger()
}

func GetLogger() *slog.Logger {
	return _logger
}

func Debugf(format string, a ...any) {
	GetLogger().Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	GetLogger().Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	GetLogger().Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	GetLogger().Error(fmt.Sprintf(format, a...))
}

func initLogger() {
	_logger = slog.New(textHandler())
}

func textHandler() *slog.TextHandler {
	logFile := getLogPath()
	file, err := os.OpenFile(logFile, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return slog.NewTextHandler(file, getHandlerOptions())
}

func getLogPath() string {
	logDir := filepath.Join(util.GetProjPath(), LOG_DIR)
	valid, _ := util.ValidPath(logDir)
	if !valid {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	return filepath.Join(logDir, fmt.Sprintf("%s.log", util.GetProjName()))
}

func getHandlerOptions() *slog.HandlerOptions {
	opt := &slog.HandlerOptions {
		AddSource: true,
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.SourceKey:
				projName := util.GetProjName()
				if v, ok := a.Value.Any().(*slog.Source); ok {
					idx := strings.Index(v.File, projName)
					a.Value = slog.StringValue(v.File[idx:])
				}
			default:
			}
			return a
		},
	}
	return opt
}




