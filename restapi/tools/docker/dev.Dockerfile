FROM golang:1.23.3

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080

CMD ["air", "-c", "/app/tools/air/.air.toml"]