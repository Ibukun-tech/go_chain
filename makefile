.PHONY:run
build:
	go build -o bin ./cmd/web
run:build
	go run ./cmd/web