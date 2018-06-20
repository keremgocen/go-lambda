.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./reactioneer/reactioneer
	
build:
	GOOS=linux GOARCH=amd64 go build -o reactioneer/reactioneer ./reactioneer