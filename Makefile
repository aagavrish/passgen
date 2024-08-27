test:
	@go test -v ./...

build:
	@go build cmd/cli/*.go

clean:
	@rm main

rebuild: clean build

.PHONY: test build clean rebuild