FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY config/ ./config/

RUN go build -a -installsuffix cgo -o server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]