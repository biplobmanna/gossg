.PHONY: all build run clean test fmt vet dev

# Variables
BINARY_NAME=gossg
MAIN_PATH=./cmd/ssg/main.go
OUTPUT_DIR=dist

# The default target run when you just type 'make'
all: fmt vet test build

# Compile the Go binary
build:
	@echo "Building the SSG engine..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)

# Build and execute the generator
run: clean build
	@echo "Running the generator..."
	@./$(BINARY_NAME)

# Clean up the binary and the generated site
clean:
	@echo "Cleaning workspace..."
	@go clean
	@rm -f $(BINARY_NAME)
	@rm -rf $(OUTPUT_DIR)/*

# Run unit tests across all internal packages
test:
	@echo "Running tests..."
	@go test ./... -v

# Format the Go code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Catch common Go errors
vet:
	@echo "Vetting code..."
	@go vet ./...

# A workspace target for development
dev: run
	@echo "Site generated in /$(OUTPUT_DIR). Ready for preview."
	@# If you add a local HTTP server or a Bubble Tea interface later, 
	@# this is where you'd trigger the watch/serve mode.
