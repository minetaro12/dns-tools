FROM node:18-bookworm-slim AS web-builder
WORKDIR /work
COPY ./web/. ./
RUN corepack enable && pnpm install && pnpm build

FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /work
COPY . ./
COPY --from=web-builder /work/build ./web/build
RUN CGO_ENABLED=0 go build -o main

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8000

CMD ["/app/main"]