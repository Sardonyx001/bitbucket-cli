# Build the application
build:
    go build -o ./build/bitbucket-cli ./cmd/bitbucket-cli

# Clean build artifacts
clean:
    rm -rf build/

# Format Go code
fmt:
    gofmt -s -w .
