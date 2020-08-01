FROM golang:1.14-alpine AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o=a

FROM alpine

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build /app/a /app/a
COPY --from=build /app/leap-seconds/sample /app/leap-seconds/sample
EXPOSE 8080
WORKDIR /app
ENTRYPOINT [ "/app/a" ]
