FROM golang:1.17-alpine AS builder
LABEL stage=builder
WORKDIR /app
COPY . .

RUN go build -o proxy cmd/main.go

FROM alpine:3.6
WORKDIR /app
COPY --from=builder /app .
CMD ["/app/proxy"]