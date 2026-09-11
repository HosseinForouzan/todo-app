FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/repository/psql/migrations ./repository/psql/migrations


EXPOSE 8080

CMD ["./server"]