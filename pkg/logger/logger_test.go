package logger_test

import (
	// "log/slog"
	"testing"
	"time"

	"github.com/easybyr/go-line/pkg/logger"
)

func TestLogFile(t *testing.T) {
	logger.GetLogger().Debug("test debug!")
	logger.GetLogger().Info("test logger!")
	logger.GetLogger().Warn("test warn!")
	logger.GetLogger().Error("test error!", "createdAt", time.Now().String())
}

