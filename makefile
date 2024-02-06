.PHONY:run
build:
	go build -o bin ./cmd/web
run:build
	go run ./cmd/web
dock:
	docker-compose up
dock-down:
	docker-compose down