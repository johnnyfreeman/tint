#!/bin/bash

# Build the application
go mod tidy
go build -o api-client

echo "Build complete. Run ./api-client to start the application."