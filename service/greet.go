package service

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

func NewGreetService(l *zap.SugaredLogger) GreetService {
	return func(names ...string) (string, error) {
		if len(names) > 5 {
			l.Debug("this place is crowded")
			return "", fmt.Errorf("that's too much persons to greet")
		}
		if len(names) == 0 {
			l.Debug("noone? let's greet anyway!")
			return "Hello, world!", nil
		}
		return fmt.Sprintf("Hello, %s", strings.Join(names, ", ")), nil
	}
}
