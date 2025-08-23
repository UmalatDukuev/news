FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o app

RUN chmod +x app

CMD ["./app"]
