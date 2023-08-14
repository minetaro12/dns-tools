FROM node:18-alpine3.18 AS web-builder
WORKDIR /work
COPY ./web/. ./
RUN corepack enable && pnpm install && pnpm build

FROM golang:1.21-alpine3.18 AS builder
WORKDIR /work
COPY . ./
COPY --from=web-builder /work/build ./web/build
RUN CGO_ENABLED=0 go build -o main

FROM alpine:3.18.3
RUN apk add --no-cache bind-tools
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8000

CMD ["/app/main"]