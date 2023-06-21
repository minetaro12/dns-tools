FROM golang:1.20.5-alpine3.18 AS builder

WORKDIR /work
COPY . ./
RUN CGO_ENABLED=0 go build -o main

FROM alpine:3.18.2
RUN apk add --no-cache bind-tools
WORKDIR /app
COPY --from=builder /work/main /app

EXPOSE 8000

CMD ["/app/main"]