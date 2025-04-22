# Create necessary directories if they don't exist
$null = New-Item -ItemType Directory -Force -Path internal/domain/models
$null = New-Item -ItemType Directory -Force -Path internal/adapters/api
$null = New-Item -ItemType Directory -Force -Path internal/adapters/client

Write-Host "Generating API models..."
oapi-codegen --config=.codegen/models.yaml ./palaestra-api-v1.yaml

Write-Host "Generating API server..."
oapi-codegen --config=.codegen/server.yaml ./palaestra-api-v1.yaml

Write-Host "Generating API client..."
oapi-codegen --config=.codegen/client.yaml ./palaestra-api-v1.yaml

Write-Host "Generation complete!"