FROM golang:1.16 AS development
WORKDIR /app
COPY go.mod go.sum./
RUN go mod download