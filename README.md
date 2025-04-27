# About Palaestra

A good fitness app.

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
3. [Project Structure](#project-structure)
4. [Using oapi-codegen with Go Fiber](#using-oapi-codegen-with-go-fiber)
5. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
6. [Development Workflow](#development-workflow)
7. [Authentication](#authentication)
8. [Hexagonal Architecture Implementation](#hexagonal-architecture-implementation)
   - [Core Components](#core-components)
   - [Data Flow](#data-flow)
9. [Code Generation](#code-generation)
   - [Using oapi-codegen](#using-oapi-codegen)
   - [Generating Code](#generating-code)
10. [Docker Support](#docker-support)
    - [Docker Setup](#docker-setup)
    - [Building and Running with Docker](#building-and-running-with-docker)
    - [Production Considerations](#production-considerations)
11. [Testing the API](#testing-the-api)
12. [Future Enhancements (TODO)](#future-enchancements-todo)

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

   or

   ```powershell
   ./scripts/generate.ps1
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

### Makefile usage

You can also run the `Makefile` for development workflow tasks.

On Windows:
```powershell
mingw32-make run
```
On Unix:
```shell
make run
```

## Authentication

The API uses OAuth2 authorization. All endpoints are secured and require an access token with appropriate scopes.

## Hexagonal Architecture Implementation

This project implements the hexagonal architecture (ports and adapters) pattern, which provides clear separation of concerns:

### Core Components

1. **Domain** (`internal/domain/models`): Contains the core business entities generated from the OpenAPI schema.

   - These models represent the fundamental data structures and business objects.
   - Generated using oapi-codegen from the OpenAPI specification.

2. **Ports** (`internal/core/ports`): Defines interfaces for interactions with the core application.
   - `repository_ports.go`: Interfaces for data persistence operations
   - `service_ports.go`: Interfaces for business logic services
3. **Services** (`internal/core/services`): Implements the business logic.

   - Contains implementation of services that connect repositories to the API.
   - Handles validation, transformation, and business rules.

4. **Adapters** (`internal/adapters`): Connects the application to external concerns.
   - `api`: HTTP handlers using Fiber, generated from OpenAPI specification
   - `repository`: Data storage implementations (currently in-memory)

### Data Flow

The hexagonal architecture enforces a clean flow of dependencies:

- External adapters → Ports → Core domain and services
- Dependencies point inward, with the domain at the center
- Core business logic doesn't depend on external frameworks or technologies

## Code Generation

The project leverages automatic code generation from the OpenAPI specification:

### Using oapi-codegen

We use oapi-codegen to generate:

1. **Domain models**: Type definitions for all API objects
2. **Server interfaces**: Required interfaces for implementing the API
3. **Fiber server**: Integration with the Fiber web framework

### Generating Code

To regenerate code after changing the OpenAPI specification:

```powershell
# For PowerShell
.\scripts\generate.ps1
```

This script:

Generates domain models in `internal/domain/models/models.gen.go`

Generates server interfaces `in internal/adapters/api/server.gen.go`

## Docker Support

The application includes Docker support for consistent deployment:

### Docker Setup

1. **Dockerfile**: Multi-stage build that compiles the Go application and creates a minimal runtime image.

2. **docker-compose.yml**: Orchestrates the application deployment with proper port mapping.

### Building and Running with Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or using Docker directly
docker build -t palaestra-api .
docker run -p 8080:8080 palaestra-api
```

### Production Considerations

For production deployment:

- The Docker image is optimized for size and security
- Uses Alpine Linux for a minimal footprint
- Proper port exposure for API access
- Configurable through environment variables

## Testing the API

### PowerShell Commands

```powershell
# Create a user
Invoke-RestMethod -Method POST -Uri "http://localhost:8080/users" -ContentType "application/json" -Body '{"email":"test@example.com","username":"testuser","firstName":"Test","lastName":"User"}'

# Get a user (replace USER_ID with actual ID)
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/users/USER_ID"

# Update a user (replace USER_ID with actual ID)
Invoke-RestMethod -Method PUT -Uri "http://localhost:8080/users/USER_ID" -ContentType "application/json" -Body '{"email":"updated@example.com","username":"updateduser","firstName":"Updated","lastName":"User"}'
```

## Future Enchancements (TODO)

Planned improvements for the project:

1. **Database Integration**: Replace in-memory repositories with actual database adapters
2. **Authentication**: Implement the OAuth2 flow described in the API specification with middleware
3. **Logging and Monitoring**: Add structured logging and metrics collection
4. **CI/CD Pipeline**: Automate testing, building, and deployment (Github workflows)
5. **Remaining Endpoints**: Complete implementation of all API endpoints defined in the specification
