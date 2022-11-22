package main

import (
	"github.com/alecthomas/kong"
	"github.com/juwit/nauclerus/cli"
)

var args cli.Args

func main() {
	ctx := kong.Parse(
		&args,
		kong.Name("nauclerus"),
		kong.Description("GUI app for managing and running Helm releases on Kubernetes"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree: true,
		}),
	)
	ctx.FatalIfErrorf(ctx.Run(&cli.Context{Verbose: args.Verbose}))
}
