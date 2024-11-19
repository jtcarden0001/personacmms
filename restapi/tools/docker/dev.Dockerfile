FROM golang:1.23.3

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app
CMD ["air", "-c", "/app/tools/air/.air.toml"]