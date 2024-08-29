OUTPUT_FILENAME=passgen

test:
	@go test -v ./...

build:
	@go build -o $(OUTPUT_FILENAME) cmd/cli/*.go
	@test -f $(OUTPUT_FILENAME) && echo "Success build"

clean:
	@rm $(OUTPUT_FILENAME)

rebuild: clean build

workflow: build clean test

.PHONY: test build clean rebuild workflow