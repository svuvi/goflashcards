.PHONY: build

BINARY_NAME=goflashcards

# build builds the tailwind css sheet, and compiles the binary into a usable thing.
build:
	templ generate && \
	go mod tidy && \
	go generate && \
	go build -ldflags="-w -s" -o ${BINARY_NAME}-new

dev:
	air

clean:
	go clean