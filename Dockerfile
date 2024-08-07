FROM node:20-bookworm-slim AS web-builder
WORKDIR /work
COPY ./web/. ./
RUN corepack enable && pnpm install && pnpm build

FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /work
COPY . ./
COPY --from=web-builder /work/build /work/web/build
RUN CGO_ENABLED=0 go build -o main

FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8000

CMD ["/app/main"]
