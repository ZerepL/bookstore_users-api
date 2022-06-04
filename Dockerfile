FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /bookstore_users-api

FROM alpine:3.13.6

COPY --from=builder /bookstore_users-api /bookstore_users-api

EXPOSE 8080

CMD ["/bookstore_users-api"]