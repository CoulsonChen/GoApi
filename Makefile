all: postgres run

postgres:
	./DockerCompose/docker-compose.yml up -d; sleep 5

build: 
	go run github.com/google/wire/cmd/wire; swag init

run: 
	go run .

down:
	docker-compose down