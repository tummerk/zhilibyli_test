FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o wallet ./cmd


FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client && \
    wget https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/wallet .
COPY db/migrations ./migrations/


ENTRYPOINT ["sh", "-c", "while ! pg_isready -h db -U postgres; do echo '⏳ Waiting for DB...'; sleep 2; done && \
    migrate -path ./migrations -database \"$DATABASE_URL\" up && \
    echo ' Migrations done. Starting app...' && \
    ./wallet"]