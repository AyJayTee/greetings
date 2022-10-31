http:
	@go env -w GOOS=linux
	@go env -w CGO_ENABLED=0
	@go build -o bin/http cmd/http/main.go
	@echo Built binary

cli: # Use to produce CLI binary
	@go env -w GOOS=linux
	@go env -w CGO_ENABLED=0
	@go build -o bin/cli cmd/cli/main.go
	@echo Built binary

image: http
	docker build . -t greetings

container: image
	@docker run -d --rm -p 8080:8080 --name hello greetings

stop:
	@docker container stop hello

reset-env:
	@go env -w GOOS=windows
	@go env -w CGO_ENABLED=1
	@echo reset go env