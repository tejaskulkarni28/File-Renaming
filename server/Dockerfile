# Build stage
FROM golang:1.17-alpine AS build
WORKDIR /app
COPY server/ ./
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
