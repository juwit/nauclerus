package nauclerus

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger(verbose bool) *zap.SugaredLogger {
	if verbose {
		l, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("cannot initialize development logger: %v", err)
		}
		defer deferredSync(l)
		return l.Sugar()
	}

	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize production logger: %v", err)
	}
	defer deferredSync(l)
	return l.Sugar()
}

func deferredSync(l *zap.Logger) {
	if err := l.Sync(); err != nil {
		log.Printf("cannot properly sync logger, some output may be missing: %s", err)
	}
}
