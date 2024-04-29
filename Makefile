migration_up:
	migrate -path db/migrations/ -database "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable" -verbose up

migration_down:
	@read -p "Enter VERSION: " VERSION; \
	migrate -path db/migrations -database "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable" -verbose down $$VERSION

migration_fix:
	@read -p "Enter VERSION: " VERSION; \
	migrate -path db/migrations -database "postgresql://postgres:@localhost:5432/cats_social?sslmode=disable" force $$VERSION