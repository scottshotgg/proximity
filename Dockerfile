FROM golang:1.15.6-alpine3.12 AS builder

WORKDIR /app

COPY . .

RUN go build -o dist/ pkg/tcphw/tcphw.go

FROM alpine:3.12

COPY --from=builder /app/dist/tcphw /bin/app

CMD ["/app/proximity", "--server", "--addrs=:9090"]