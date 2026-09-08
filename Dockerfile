# Stage 1: Build binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY *.go ./

# Build static binary without CGO
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Copy static assets and excel data
COPY static ./static
COPY export_v2.xlsx .
COPY export.xlsx .

# Ensure uploads directory exists
RUN mkdir -p uploads

EXPOSE 80

ENV PORT=80
ENV EXCEL_FILE=export_v2.xlsx

ENTRYPOINT ["./main"]
