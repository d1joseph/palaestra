# About Palaestra

A really good ftness app.

This repository contains:

- Backend of [palaestra.fit](https://palaestra.fit)
- Infrastructure as code
- CICD Workflows
- Testing and performance bench marking utilities
- Build, run and test locally before deploying

## Contents

1. [Features](#features)
2. [Requirements](#requirements)
   - [Development](#development)
3. [Project Structure](#prerequisites)

## Features

- Open API compliant REST API
- Code generation using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)

## Requirements

### Development

- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) (only binary installations are supported for contribututions)

## Project Structure

This project follows the hexagonal architecture pattern (also known as ports and adapters):

```
palaestra-api/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── adapters/             # Implementation of ports (adapters)
│   │   ├── api/              # HTTP API adapter (Fiber)
│   │   │   ├── handler.go    # API handler implementation
│   │   │   ├── middleware.go # API middleware
│   │   │   ├── router.go     # Route definitions
│   │   │   └── server.gen.go # Generated OpenAPI code
│   │   ├── client/           # Generated client
│   │   │   └── client.gen.go
│   │   └── repository/       # Data storage implementation
│   │       ├── user_repository.go
│   │       ├── exercise_repository.go
│   │       ├── workout_repository.go
│   │       ├── workout_plan_repository.go
│   │       └── progress_metric_repository.go
│   ├── core/
│   │   ├── ports/            # Interface definitions
│   │   │   ├── repository_ports.go
│   │   │   └── service_ports.go
│   │   └── services/         # Business logic
│   │       ├── user_service.go
│   │       ├── exercise_service.go
│   │       ├── workout_service.go
│   │       ├── workout_plan_service.go
│   │       └── progress_metric_service.go
│   └── domain/
│       └── models/           # Domain models
│           └── models.gen.go # Generated from OpenAPI schemas
├── scripts/
│   └── generate.sh           # Code generation script
├── palaestra-api-v1.yml      # OpenAPI 3.0 specification
├── oapi-codegen-config.yaml  # Configuration for code generation
├── go.mod
└── go.sum
```

## Using oapi-codegen with Go Fiber

This project demonstrates how to use the [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) tool with Go Fiber for API development. Unlike the standard approach that generates a complete server implementation, this project uses oapi-codegen only for model generation while manually implementing the Fiber handlers.

### The approach consists of:

1. **Using oapi-codegen for models**: Generate Go structs from OpenAPI schemas
2. **Manual Fiber implementation**: Implement the API routes manually using Fiber

### Steps to set up a similar project:

1. **Define your OpenAPI specification** in YAML or JSON format
2. **Create a configuration file** for oapi-codegen:

   ```yaml
   # oapi-codegen-config.yaml
   server:
     package: api
     out: internal/adapters/api/server.gen.go
     generate:
       embedded-spec: true

   types:
     package: models
     out: internal/domain/models/models.gen.go
     generate:
       types: true
       skip-prune: true
   ```

3. **Generate the code**:
   ```bash
   oapi-codegen -config oapi-codegen-config.yaml -generate types -package models -o internal/domain/models/models.gen.go palaestra-api-v1.yml
   ```
4. **Implement Fiber handlers**:
   - Create handler implementations in `internal/adapters/api/handler.go`
   - Define routes in `internal/adapters/api/router.go`
   - Set up middleware in `internal/adapters/api/middleware.go`
5. **Wire everything together** in `cmd/api/main.go`

## Getting Started

### Prerequisites

- Go 1.18 or higher
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)

### Installation

1. Clone the repository

   ```bash
   git clone https://github.com/yourusername/palaestra-api.git
   cd palaestra-api
   ```

2. Install dependencies

   ```bash
   go mod download
   ```

3. Generate code from OpenAPI spec

   ```bash
   ./scripts/generate.sh
   ```

4. Build and run the application
   ```bash
   go build -o palaestra-api ./cmd/api
   ./palaestra-api
   ```

The API will be available at http://localhost:8080/v1

## Development Workflow

1. Modify the OpenAPI specification (`palaestra-api-v1.yml`)
2. Run the code generation script to update models and embedded spec
   ```bash
   ./scripts/generate.sh
   ```
3. Update your handler implementations as needed
4. Build and run the application

## Authentication

The API uses OAuth2 authorization. All endpoints are secured and require an access token with appropriate scopes.
