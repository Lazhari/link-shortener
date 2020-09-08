# Go micro-service using Hexagonal Architecture

> Link shortener is an implementation of the hexagonal architecture using Go, MongoDB and Redis.

![Hexagonal Architecture](doc/hexagonal-architecture.png)

## Starting the service

### Redis

```
export URL_DB=redis
export REDIS_URL=redis://localhost:6379
go run .
```

### MongoDB

```
export URL_DB=mongo
export MONGO_URL=mongodb://localhost/shortener
export MONGO_DB=shortener
export MONGO_TIMEOUT=30
go run .
```
