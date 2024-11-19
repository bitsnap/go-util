package goutil

import (
	"log"
	"os"
	"strings"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func SetGlobalLogger() func() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}

	if strings.ToLower(os.Getenv("NODE_ENV")) == "production" {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalln(err)
		}
	}

	globalOtelLogger := otelzap.New(logger)
	undo := otelzap.ReplaceGlobals(globalOtelLogger)

	return func() {
		_ = globalOtelLogger.Sync()
		undo()
	}
}

func Logger() *otelzap.SugaredLogger {
	return otelzap.L().Sugar()
}
