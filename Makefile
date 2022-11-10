http:
	@go env -w GOOS=linux
	@go env -w CGO_ENABLED=0
	@go build -o bin/http cmd/http/main.go
	@echo Built binary
.PHONY: http

cli: # Use to produce CLI binary
	@go env -w GOOS=linux
	@go env -w CGO_ENABLED=0
	@go build -o bin/cli cmd/cli/main.go
	@echo Built binary
.PHONY: cli

image: http
	docker build . -t greetings
.PHONY: http

container: image
	@docker run -d --rm -p 8080:8080 --name hello greetings
.PHONY: container

start:
	@env GOOS=darwin GOARCH=amd64 go run ./cmd/http/main.go
.PHONY: start

dev:
	@find . \( -path './vendor/*' -or -path './assets/*' \) -prune -false -o -name '*.go' \
	| entr -rcs 'make start'
.PHONY: dev

stop:
	@docker container stop hello
.PHONY: stop

reset-env:
	@go env -w GOOS=windows
	@go env -w CGO_ENABLED=1
	@echo reset go env
.PHONY: reset-env
