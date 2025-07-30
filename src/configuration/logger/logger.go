package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log *zap.Logger

	LOG_OUTPUT = os.Getenv("LOG_OUTPUT")
	LOG_LEVEL  = os.Getenv("LOG_LEVEL")
)

func init() {

	logConfig := zap.Config{
		OutputPaths: []string{getOutputLog()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			// Define the keys for the log fields
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func getOutputLog() string {
	output := strings.ToLower(strings.TrimSpace(LOG_OUTPUT))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLevelLogs() zapcore.Level {
	level := strings.ToLower(strings.TrimSpace(LOG_LEVEL))
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
