SWAG := swag
MOCKERY := mockery
APP_PATH := $(APP_DIR)/cmd/server/main.go


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


.PHONY: dev dev-worker-a mocks docs