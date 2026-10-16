FROM node:24-trixie-slim AS web-builder
WORKDIR /work
COPY ./web/. ./
RUN corepack enable pnpm && pnpm install && pnpm run build

FROM golang:1.27.2-alpine3.24 AS builder
WORKDIR /work
COPY . ./
COPY --from=web-builder /work/dist /work/web/dist
RUN CGO_ENABLED=0 go build -o main

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8080

CMD ["/app/main"]
