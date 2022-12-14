package application

import (
	"github.com/juwit/nauclerus"
	"github.com/juwit/nauclerus/http"
	"github.com/juwit/nauclerus/service"
)

// Build basic dependency injection
func Build(verbose, debugMode bool) (application Application, err error) {
	logger := nauclerus.InitLogger(verbose)
	logger.Debugf("debug output is on")

	// Instantiate beans here
	greetService := service.NewGreetService(logger.With("layer", "service"))
	greetMiddleware := http.NewGreetMiddleware(greetService, logger.With("layer", "http"))

	router := http.NewRouter(greetMiddleware, debugMode, logger.With("layer", "routing"))

	application.Logger = logger
	application.DebugMode = debugMode
	application.Router = router
	return
}

// A deferable method of things we will have to do to close the application (shut down DB, etc...)
func (app Application) Close() error {
	return nil
}
