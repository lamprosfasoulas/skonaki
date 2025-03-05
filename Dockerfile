FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

RUN apk add --no-cache bat aha

RUN addgroup -S -g 1500 skon && adduser -S -u 1500 skon -G skon

USER skon

WORKDIR /app

COPY --chown=skon:skon --from=builder /app /app

EXPOSE 42069

CMD ["./main"]
