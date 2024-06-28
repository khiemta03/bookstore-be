FROM golang:1.22.4-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
RUN apk add --no-cache postgresql-client libc6-compat
ENV LANG en_US.utf8
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 /app/migrate
COPY script /app/script
COPY configs /app/configs
COPY migration /app/migration

EXPOSE 3000
CMD ["/app/main"]
ENTRYPOINT [ "/app/script/start.sh" ]

