dev:
	@docker-compose up -d
	@GO_ENV=dev  go run main.go