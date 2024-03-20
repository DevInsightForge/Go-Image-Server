# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
ENV ENVIRONMENT production

COPY go.mod go.sum /app/
RUN set -Eeux && \
    go mod download && \
    go mod verify

COPY . /app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -o image-server cmd/api/main.go

# Final stage
FROM scratch
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/image-server /app/main

ENV ADDRESS=0.0.0.0
ENV PORT 4000
EXPOSE 4000

CMD ["./main"]
