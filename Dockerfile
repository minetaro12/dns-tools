FROM golang:1.20.1-alpine3.17 AS builder

WORKDIR /work
COPY . ./
RUN CGO_ENABLED=0 go build

FROM alpine:3.17.2
RUN apk add --no-cache bind-tools
WORKDIR /app
COPY --from=builder /work/dns-tools /app

EXPOSE 8000

CMD ["/app/dns-tools"]