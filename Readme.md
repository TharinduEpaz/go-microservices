# Go Microservices with Redis

This is a simple example of how to implement microservices with Redis in Go.

## Features

- Redis storage for the data
- Graceful shutdown

## Redis setup using Docker

```bash
docker run -p 6379:6379 redis:latest
```

## Run the server

```bash
go run main.go
```
