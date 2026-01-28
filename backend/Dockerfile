FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    && rm -rf /var/lib/apt/lists/*

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./main"]
