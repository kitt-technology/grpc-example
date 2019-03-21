FROM golang:1.11.1-alpine3.8

RUN apk add --no-cache \
    ca-certificates \
    git \
    gcc \
    openssh

COPY . /app

WORKDIR /app/src

ENV CGO_ENABLED 0
ENV GOOS linux

RUN GO111MODULE=auto go mod vendor

CMD ["go", "run", "main.go"]
