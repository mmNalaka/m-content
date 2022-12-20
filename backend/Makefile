# Databse url
DB_URL = postgres://postgres:postgres@localhost:5432/m_content_development?sslmode=disable
PATH_TO_SCHEMA = db/migrations/

# SQLC commands
sqlc:
	sqlc generate

# Migration commands
migrate.create:
	migrate create -ext sql -dir $(PATH_TO_SCHEMA) -seq $(name)

migrate.up:
	migrate -database $(DB_URL) -path $(PATH_TO_SCHEMA) up

migrate.down:
	migrate -database $(DB_URL) -path $(PATH_TO_SCHEMA) down

# Docker commands
docker.postgres: 
	docker run --name m_content_development -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=m_content_development -d postgres:14-alpine