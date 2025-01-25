FROM golang:1.21.6-alpine3.18 as Build

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /ping_app main.go

FROM alpine:latest

WORKDIR /
COPY --from=Build /ping_app /ping_app

ENTRYPOINT [ "/ping_app" ]