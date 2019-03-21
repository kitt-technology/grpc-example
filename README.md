# Minimal Go + gRPC Example

This repository provides a minimal example for communication via gRPC between two Golang microservices. Service A requests a message from Service B, Service B responds with “pong”.

## Installation

To get things running:
```
docker-compose up
```
Navigate to [http://localhost:8080/ping](http://localhost:8080/ping) and see “pong” being returned from a Service B.