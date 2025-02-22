FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

RUN apk add --no-cache bat aha

WORKDIR /app

COPY --from=builder /app .

EXPOSE 42069

CMD ["./main"]
