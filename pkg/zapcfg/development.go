package zapcfg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDevelopment returns a new zap logger config for development environment.
func NewDevelopment(debugMode bool, outputPath string) zap.Config {
	var logLevel zapcore.Level
	if debugMode {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.InfoLevel
	}

	return zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: true,
		Encoding:    "console",

		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},

		OutputPaths:      []string{outputPath},
		ErrorOutputPaths: []string{outputPath},
	}
}
