FROM golang:1.22.11-alpine3.21 as Build

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /pong_app main.go

FROM alpine:latest

WORKDIR /
COPY --from=Build /pong_app /pong_app

ENTRYPOINT [ "/pong_app" ]