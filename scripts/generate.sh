#!/bin/bash

# Ensure the script exits on error
set -e

# Create necessary directories
mkdir -p internal/adapters/api
mkdir -p internal/domain/models
mkdir -p internal/adapters/client

# Generate types (models) code
echo "Generating model types..."
oapi-codegen -config oapi-codegen-config.yaml -generate types -package models -o internal/domain/models/models.gen.go palaestra-api-v1.yaml

# Generate embedded spec 
echo "Generating embedded spec..."
oapi-codegen -config oapi-codegen-config.yaml -generate embedded-spec -package api -o internal/adapters/api/spec.gen.go palaestra-api-v1.yaml

# Generate client code (optional)
echo "Generating client code..."
oapi-codegen -config oapi-codegen-config.yaml -generate client -package client -o internal/adapters/client/client.gen.go palaestra-api-v1.yaml

# Make sure the script is executable
chmod +x scripts/generate.sh

echo "Code generation complete!"