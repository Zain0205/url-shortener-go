# ===============================
# STAGE 1: Build the Go binary
# ===============================
FROM golang:1.25.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ===============================
# STAGE 2: Run the app
# ===============================
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 3000
CMD ["./main"]

