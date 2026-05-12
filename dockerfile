FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o issuetracker .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/issuetracker .
COPY migrations/ ./migrations/
EXPOSE 8080
CMD ["./issuetracker"]