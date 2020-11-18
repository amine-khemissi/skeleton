build:
	@CGO_ENABLED=0 go build -o bin/svc main/main.go

up: build
	@docker-compose   up  --remove-orphans    --build web

stop:
	@docker-compose stop
