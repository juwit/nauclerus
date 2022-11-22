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
go run cmd/nauclerus/main.go
```
