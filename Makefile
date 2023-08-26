build:
	@cd web && pnpm i && pnpm build
	@go build -o main

dev:
	@cd web && pnpm i && pnpm build
	@go run .

web-dev:
	@cd web && pnpm i && pnpm dev