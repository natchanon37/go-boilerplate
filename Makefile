SWAG := swag
MOCKERY := mockery


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


.PHONY: dev dev-worker-a mocks