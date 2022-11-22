package cli

import (
	"fmt"

	"github.com/juwit/nauclerus/application"
)

func (c ServeCmd) Run(context *Context) error {
	app, err := application.Build(context.Verbose, c.Debug)
	if err != nil {
		return err
	}
	defer func() {
		_ = app.Close()
	}()
	return app.Router.Run(fmt.Sprintf(":%d", c.Port))
}
