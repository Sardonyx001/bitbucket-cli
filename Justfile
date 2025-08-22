# Build the application
build:
    mkdir -p build/
    go build -o ./build/bitbucket-cli ./cmd/bitbucket-cli

# Clean build artifacts
clean:
    rm -rf build/