image:
	docker build . -t greetings
.PHONY: image

container: image
	@docker run -d --rm -p 8080:8080 --name hello greetings
.PHONY: container

start:
	@go run ./cmd/http/main.go
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
