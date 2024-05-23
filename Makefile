.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: build
build: ## BUILD - Build the application
	CGO_ENABLED=0 GOOS=linux go build -a -o build/clearing main.go

.PHONY: vuln
vuln: ## VULN - Check for vulnerabilities
	# check for vulnerabilities...
	@govulncheck ./...

.PHONY: migrate-up
migrate-up: ## Run the migrations
	# Run the migrations
	# ------------------
	@migrate -database="postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE)" -path=migrations -verbose up

.PHONY: migrate-down
migrate-down: ## Rollback the last migration
	# Rollback the last migration
	# ----------------------------
	@migrate -database="postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE)" -path=migrations -verbose down

.PHONY: migrate-create
migrate-create: ## Create a new migration file. ex: make migrate-create name=create_users_table
	# Create a new migration file
	# ----------------------------
	@migrate create -ext sql -dir=migrations -seq $(name)

.PHONY: migrate-drop
migrate-drop:
	# Drop all tables
	# ---------------
	@migrate -database="postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE)" -path=migrations -verbose drop -f

.PHONY: migrate-force
migrate-force: 
	@migrate -database="postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE)" -path=migrations -verbose force 1