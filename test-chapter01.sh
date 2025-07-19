#!/bin/bash
set -e

echo "Starting test servers for Chapter 01..."

# Start go-httpbin on port 8080
docker run -d --rm -p 8080:8080 --name go-httpbin mccutchen/go-httpbin

# Build and start custom test server
cd examples/chapter01/test-server
go build -o test-server main.go
./test-server &
TEST_SERVER_PID=$!
cd ../../..

# Wait for servers to start
echo "Waiting for servers to start..."
sleep 3

# Function to cleanup on exit
cleanup() {
    echo "Cleaning up..."
    docker stop go-httpbin || true
    kill $TEST_SERVER_PID 2>/dev/null || true
}
trap cleanup EXIT

# Run tests
echo "Running Chapter 01 tests..."
runn run examples/chapter01/**/*.yml

echo "All Chapter 01 tests passed!"