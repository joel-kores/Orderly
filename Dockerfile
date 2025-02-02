FROM golang:1.21-alpine AS builder
LABEL authors="Joel Kores"

ENTRYPOINT ["top", "-b"]

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o my-go-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/my-go-app .

EXPOSE 8080

CMD ["./my-go-app"]