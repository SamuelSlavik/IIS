# syntax=docker/dockerfile:1

FROM golang:latest as dev

WORKDIR /api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

CMD ["air"]
