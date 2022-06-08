# syntax=docker/dockerfile:1 
FROM golang

FROM golang:1.16

WORKDIR C:\Users\joeo1\Desktop\Go\discord_bot\main.go
COPY . .

RUN go install -v ./...
RUN go run main.go