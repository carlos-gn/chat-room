# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /api .

EXPOSE 3000
ENTRYPOINT ["/api"]
