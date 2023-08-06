include .env
export

BASE_DIR := $(shell pwd)
CREATE_MIGRATION_SCRIPT := $(BASE_DIR)/scripts/create_migrations.sh
RUN_MIGRATIONS_SCRIPT := $(BASE_DIR)/scripts/run_migrations.sh

create_migration:
ifndef NAME
	@echo "Error: NAME parameter not provided."
	@exit 1
endif

	@echo "Creating migration script..."
	@bash $(CREATE_MIGRATION_SCRIPT) -d "internal/db/migrations/" -n $(NAME)


migrate:
ifndef ARG
	@echo "Error: ARG parameter not provided."
	@exit 1
endif

	@echo "Running migrations..."
	@bash $(RUN_MIGRATIONS_SCRIPT) -$(ARG)

generate_sql:
	@echo "Generating SQL..."
	@docker run --rm -v $(BASE_DIR)/internal/db:/src -w /src kjconroy/sqlc generate
