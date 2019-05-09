build:
	@docker build . -t statbot

start:
	@docker run -it -rm -name statbot statbot