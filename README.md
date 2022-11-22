# Nauclerus

## How to build

The preferred way to build the application is either using the Makefile's `build` target.

You should have [`golangci-lint` installed](https://golangci-lint.run/usage/install/) in order to perform various quality checks.

The build has only been tested on go 1.19.

You can see every build targets simply running `make` with arguments. To generate the binary, run

```
make build
```

## How to run

If you do not need to generate a binary, you can simply `go run` (or `make run`) the program and see the possible options.

```
$ go run cmd/nauclerus/main.go serve --help

Usage: nauclerus serve

Start the server

Flags:
-h, --help         Show context-sensitive help.
-v, --verbose      Enable debug output ($NAUCLERUS_VERBOSE)

--debug        weither to enable debug mode ($NAUCLERUS_DEBUG)
-p, --port=8080    HTTP Port to listen to ($NAUCLERUS_PORT)
```

### Configuration

Run the app with --help to see configuration options. Most configuration controls are available as environment variables (useful in a docker context)

## Architecture

The project layout can be documented like this:

```
.
├── application             # application container
├── cli                     # CLI controllers, process command line arguments and start the program
├── cmd                     # main func entrypoints, not go gettable
├── http                    # HTTP controllers, process HTTP logic and routing
├── service                 # Business logic for this service
└── logger.go               # logging utils
```