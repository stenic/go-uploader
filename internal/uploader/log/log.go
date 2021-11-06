package log

import (
	"context"
	"os"

	logrus "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

const loggerKey = "logger"

func ContextWithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey)

	if logger == nil {
		return logrus.WithContext(ctx)
	}

	return logger.(*logrus.Entry)
}
