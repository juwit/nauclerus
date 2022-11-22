package cli

type Context struct {
	Verbose bool
}

type ServeCmd struct {
	Debug bool `help:"weither to enable debug mode" env:"NAUCLERUS_DEBUG"`
	Port  int  `help:"HTTP Port to listen to" env:"NAUCLERUS_PORT" short:"p" default:"8080"`
}

type Args struct {
	Verbose bool `help:"Enable debug output" env:"NAUCLERUS_VERBOSE" short:"v"`

	Serve ServeCmd `cmd:"" help:"Starts the server"`
}
