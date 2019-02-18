.PHONY: clean build

clean: 
	rm -rf ./hello/hello-world
	
build:
	
	GOOS=linux GOARCH=amd64 go build -o code/hello-world ./code/src/hello
	GOOS=linux GOARCH=amd64 go build -o code/callback ./code/src/callback
	GOOS=linux GOARCH=amd64 go build -o code/greet ./code/src/greet
	GOOS=linux GOARCH=amd64 go build -o code/initApi ./code/src/initApi
