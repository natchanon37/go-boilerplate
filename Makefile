SWAG := swag
MOCKERY := mockery
APP_PATH := $(APP_DIR)/cmd/server/main.go
ATLATGO := atlas

MYSQL_HOST := 127.0.0.1
MYSQL_PORT := 4406
MYSQL_ROOT_PASSWORD := 111
MYSQL_ROOT_USER := root
MYSQL_DATABASE := mydb


# Generate swagger documentation
docs:
	@echo "📚 Generating swagger documentation"
	@$(SWAG) init --parseDependency -g $(APP_PATH)
# Generate test mocks
mocks:
	@echo "📚 Generating test mocks"
	@$(MOCKERY) --all

# Run dev
dev-worker-a:
	@echo "💻 Running worker a kafka server..."
	@air -c .air-worker-a.toml
dev:
	@echo "💻 Running dev server..."
	@air .


# Database apply schema to local
db-apply:
	@echo "💻 apply migrations to local"
	@$(ATLATGO) schema apply --env gorm -u "mysql://$(MYSQL_ROOT_USER):$(MYSQL_ROOT_PASSWORD)@$(MYSQL_HOST):$(MYSQL_PORT)/$(MYSQL_DATABASE)"
	@echo "✅ DONE apply migration to local\n"


# Database diff schema from model
db-diff:
	@echo "💻 generate migration files"
	@$(ATLATGO) migrate diff --env gorm


.PHONY: dev dev-worker-a mocks docs db-apply db-diff