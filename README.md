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
