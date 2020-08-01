# mars-time
Go micro-service to convert UTC Earth time to Mars time

## HTTP API

Server listens on port 8080 and exposes the following endpoint:

### POST `/api/v1/convert`

Required header: `Content-Type: application/json`
Expected body: 
```json
{
  "earth": "2020-01-01T00:00:00Z"
}
```
Date should be in the RFC3339 format.

## Docker image

Use the Docker image to test locally or deploy to an environment.

### Build

`docker build -t mars-time .`

### Run

`docker run --rm -ti -p 8080:8080 mars-time`

## Local development

1. Make sure Go is installed
2. `go test ./...` to run the test suite
