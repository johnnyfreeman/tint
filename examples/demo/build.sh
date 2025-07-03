#!/bin/bash

# Build the demo application
go mod tidy
go build -o demo

echo "Build complete. Run ./demo to start the application."