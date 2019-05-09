build:
	@docker build . -t statbot

start:
	@docker build . -t statbot
	@docker run --rm --env-file .env --name statbot statbot

dev:
	@go run main.go