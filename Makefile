.PHONY: build clean deploy run

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/drnkbot main.go message_creator.go

clean:
	rm -rf ./bin

deploy: clean build
	bash -c "set -a && source .env.production && set +a && sls deploy --verbose"

run: clean build
	bash -c "set -a && source .env && set +a && bin/drnkbot"
