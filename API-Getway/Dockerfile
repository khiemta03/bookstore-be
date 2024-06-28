FROM golang:1.22.4-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY configs /app/configs
COPY public /app/public

EXPOSE 3000
CMD ["/app/main"]