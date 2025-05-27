package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func GetLogger(env string) *zap.Logger {

	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " | ",
	})

	var logLevel zapcore.Level

	switch env {
	case "local":
		logLevel = zap.DebugLevel
	case "dev":
		logLevel = zap.DebugLevel
	case "prod":
		logLevel = zap.InfoLevel
	}

	consoleWriter := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(consoleEncoder, consoleWriter, logLevel)
	logger := zap.New(core, zap.AddCaller())

	return logger

}
