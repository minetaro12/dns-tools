FROM node:18-alpine3.18 AS web-builder
WORKDIR /work
COPY ./web/. ./
RUN yarn install && yarn build

FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /work
COPY . ./
COPY --from=web-builder /work/build ./web/build
RUN CGO_ENABLED=0 go build -o main

FROM alpine:3.18.2
RUN apk add --no-cache bind-tools
WORKDIR /app
COPY --from=builder /work/main /app/main

EXPOSE 8000

CMD ["/app/main"]