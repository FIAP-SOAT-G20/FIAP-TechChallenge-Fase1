FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/http/main.go

FROM alpine:3.20.3 AS final

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
