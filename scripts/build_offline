#!/bin/bash

# Enable Go Modules.
# This ensures the build process uses the go.mod file for dependency management.
go env -w GO111MODULE=on

echo "Starting build process using vendored dependencies from Git..."

# Build the executables using the -mod=vendor flag.
# This forces the Go compiler to use the local 'vendor' directory
# and prevents any network access for dependencies.
go build -mod=vendor -o ./ecdsagen ./src/main/ecdsagen
./ecdsagen 0 100

go build -mod=vendor -o ./server ./src/main/server
go build -mod=vendor -o ./client ./src/main/client

echo "Build finished successfully!"

# List the generated binaries to confirm they were created.
ls -l ecdsagen server client
