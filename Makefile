include .env

migration_up:
	migrate -path db/migrations/ -database "postgresql://${DB_USERNAME}:@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migration_down:
	@read -p "Enter VERSION: " VERSION; \
	migrate -path db/migrations -database "postgresql://${DB_USERNAME}:@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down $$VERSION

migration_fix:
	@read -p "Enter VERSION: " VERSION; \
	migrate -path db/migrations -database "postgresql://${DB_USERNAME}:@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $$VERSION