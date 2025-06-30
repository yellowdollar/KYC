FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/swaggo/swag/cmd/swag@latest init -g main.go --output docs --outputTypes go,json,yaml

RUN go build -o main .

CMD ["./main"]