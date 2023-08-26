build:
	@cd web && pnpm i && pnpm build
	@go build -o main

dev:
	@cd web && pnpm i && pnpm build
	@go run .