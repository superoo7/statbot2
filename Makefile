init:
	@(cd nodejs && npm install)
	@GO111MODULE=on go get
	@npm i -g concurrently
	@mkdir img

build:
	@docker build . -t statbot

start:
	@docker build . -t statbot
	@docker run --rm --env-file .env --name statbot statbot

dev:
	@concurrently \
	--names "GO,JS" -c "bgBlue.bold,bgGreen.bold" \
	"go run main.go" "(cd nodejs && node index.js)"

