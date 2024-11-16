# docker build -t restapi-dev -f tools/docker/dev.Dockerfile -p 8080:8080 .
FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@v1.61.1

EXPOSE 8080

RUN ls -al /app

CMD ["air", "-c", "/app/tools/air/.air.toml"]