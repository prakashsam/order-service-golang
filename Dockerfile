FROM golang:1.24.5 as builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o order-service main.go

FROM gcr.io/distroless/static

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/order-service /

CMD ["/order-service"]
