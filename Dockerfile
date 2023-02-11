FROM golang:1.20.0 AS builder

WORKDIR /work
COPY . ./
RUN CGO_ENABLED=0 go build

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=builder /work/dns-tools /app

EXPOSE 8000

CMD ["/app/dns-tools"]