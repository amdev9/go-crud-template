FROM golang:1.21-alpine AS build_base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-crud-template .

EXPOSE 8080


ENTRYPOINT ["./go-crud-template"]
