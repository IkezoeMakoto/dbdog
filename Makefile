build:
	docker-compose up -d app
.PHONY: build

APP=./app/bin/dbdog
up:
	$(APP)
.PHONY:up