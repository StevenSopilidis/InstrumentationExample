FROM golang:1.22.11-alpine3.21 as Build

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /ping_app main.go

FROM alpine:latest

WORKDIR /
COPY --from=Build /ping_app /ping_app

ENTRYPOINT [ "/ping_app" ]