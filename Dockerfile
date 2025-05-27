# Stage 1 - build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o employee_service cmd/employee_service/main.go && \
    go build -o migrator cmd/migrator/main.go

# Stage 2 - run
FROM alpine
WORKDIR /app
COPY --from=builder /app/employee_service .
COPY --from=builder /app/migrator .
COPY ./migrations ./migrations
CMD ["./user_service"]