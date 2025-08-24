FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env /app

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# CMD ["./app"]


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

COPY --from=builder /app/config ./config

EXPOSE 8000

CMD ["./server"]