include .env
export

CREATE_MIGRATION_SCRIPT := $(shell pwd)/scripts/create_migrations.sh
RUN_MIGRATIONS_SCRIPT := $(shell pwd)/scripts/run_migrations.sh

create_migration:
ifndef DIR
	@echo "Error: DIR parameter not provided."
	@exit 1
endif
ifndef NAME
	@echo "Error: NAME parameter not provided."
	@exit 1
endif

	@echo "Creating migration script..."
	@bash $(CREATE_MIGRATION_SCRIPT) -d "internal/db/migrations/$(DIR)" -n $(NAME)


migrate:
ifndef ARG
	@echo "Error: ARG parameter not provided."
	@exit 1
endif

	@echo "Running migrations..."
	@bash $(RUN_MIGRATIONS_SCRIPT) -$(ARG)

generate_sql:
	@echo "Generating SQL..."
	@docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
