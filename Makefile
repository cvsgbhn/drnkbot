.PHONY: build clean deploy run

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/drnkbot drnkbot/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

run: clean build
	bin/drnkbot
