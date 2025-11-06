DB_PATH := bookworm.db
MIGRATIONS_DIR := migrations

.PHONY: watch migrate up down create status

# Run the project with hot reload active
watch:
	air

# Run all migrations
migrate:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up

# Run the next migration
up:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up-by-one

# Roll back the last migration
down:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) down

# Create a new migration file
# Usage: make create NAME=create_users_table
create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make create NAME=migration_name"; \
		exit 1; \
	fi
	goose create $(NAME) sql

# Status of migrations
status:
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) status

