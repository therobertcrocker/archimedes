# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Binary name
BINARY_NAME = archimedes

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/archimedes

# Clean the build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/archimedes
	./$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Install project dependencies
deps:
	$(GOGET) -v ./...

# Default make target
default: build
