# syntax=docker/dockerfile:1

FROM golang:latest as dev

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /iis-api

CMD ["/iis-api"]
