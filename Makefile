all: help

help: ## display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

run: install ## run the app
	go run cmd/nauclerus/main.go

lint: ## lint code (requires golangci-lint)
	golangci-lint run

mod-verify: ## verify dependencies
	go mod verify

test: ## run tests
	go test ./...

ci: mod-verify lint test ## run all CI checks

install: ## install dependencies
	go mod download

build: install ci ## generate binary
	go build -o nauclerus cmd/nauclerus/main.go

clean: ## clean build output
	rm -f nauclerus

.PHONY: all help lint mod-verify test ci install build clean
