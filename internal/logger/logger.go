package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Settings struct {
    // Key for the timestamp field
    TimeKey         string
    // Key for the level field
    LevelKey        string
    // Key for the logger name field
    NameKey         string
    // Key for the caller field
    CallerKey       string
    // Key for the message field
    MessageKey      string
    // Key for the stacktrace field
    StacktraceKey   string

    // Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
    Level           zapcore.Level
    // DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
    DisableCaller   bool
    // Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
    Development     bool
}

var DevelopmentSettings = Settings{
    Level:          zapcore.DebugLevel,
    DisableCaller:  true,
    Development:    true,
}

var ProductionSettings = Settings{
    Level:          zapcore.ErrorLevel,
    DisableCaller:  true,
    Development:    false,
}

var SettingNames = map[string]Settings{
    "development": DevelopmentSettings,
    "production": ProductionSettings,
}

func NewLogger(setting string) (*zap.Logger, error) {
    var err error

    s := SettingNames[setting]

    encoderConfig := zapcore.EncoderConfig{
        TimeKey: "timestamp",
        LevelKey: "level",
        NameKey: "logger",
        CallerKey: "caller",
        MessageKey: "message",
        StacktraceKey: "stacktrace",
		LineEnding: "\n",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	config := zap.Config{
		Level: zap.NewAtomicLevelAt(s.Level),
        DisableCaller: s.DisableCaller,
		Development: s.Development,
		Encoding: "json",
		EncoderConfig: encoderConfig,
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

    logger, err := config.Build()

	if err != nil {
        return nil, fmt.Errorf("logger: failed to build a new logger instance: %w", err)
	}

    return logger, nil
}
