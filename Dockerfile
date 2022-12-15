# STAGE 1: Build
FROM golang:1.19.3-alpine3.15 as builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# STAGE 2: Run
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/util/mailing/templates ./util/mailing/templates
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# install ca certificate
RUN apk update
RUN apk add ca-certificates

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]