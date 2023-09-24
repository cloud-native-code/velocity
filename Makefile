# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Output binary name
BINARY = build/launcher
GOENTRYPOINT = cmd/launcher/main.go

all: clean build

build:
	$(GOBUILD) -o $(BINARY) -v $(GOENTRYPOINT)

clean:
	$(GOCLEAN)
	rm -f $(BINARY)

run: build
	./$(BINARY)

get-deps:
	$(GOGET) .

.PHONY: all build clean run get-deps
