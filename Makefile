.PHONY: clean build deps api

deps: 
	go mod download

clean: 
	rm -rf ./bin/*
	
build:
	
	GOOS=linux GOARCH=amd64 go build -o bin/callback ./code/src/callback
	GOOS=linux GOARCH=amd64 go build -o bin/accounts ./code/src/accounts
	GOOS=linux GOARCH=amd64 go build -o bin/initApi ./code/src/initApi

api:
	make build
	sam local start-api
