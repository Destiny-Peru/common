package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func New(env string) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, fmt.Errorf("create logger: %w", err)
	}

	return logger, nil
}

func Sync(logger *zap.Logger) error {
	if logger == nil {
		return nil
	}

	return logger.Sync()
}
