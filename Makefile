run: build
	@./bin/goredis --listenAddr :8080

build:
	@go build -o bin/goredis .