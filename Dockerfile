# Gunakan image Golang untuk build
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh isi project
COPY . .

# Build binary dari cmd/main.go
RUN go build -o main ./cmd/main.go

# Runtime container
FROM debian:bookworm-slim

WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/main .

# Expose port aplikasi (ubah sesuai yang digunakan)
EXPOSE 10000

CMD ["./main"]