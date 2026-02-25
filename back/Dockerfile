# Build stage
FROM golang:1.25-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o backend-app

# Run stage
FROM alpine:latest

COPY --from=build /build/backend-app ./
EXPOSE 8080

CMD ["./backend-app"]