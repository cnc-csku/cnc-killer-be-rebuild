include .env

DB_URI="postgresql://$(DB_USERNAME):$(DB_PASSWORD)@${DB_HOST}:${DB_PORT}/$(DB_DATABASE)?sslmode=${DB_SSL_MODE}"
MIGRATIONS_PATH=internal/migrations
step=1
re-docker:
	docker compose down && docker compose up -d --build
migrate-make:
	migrate create -ext=sql -dir=$(MIGRATIONS_PATH) -tz "UTC" $(name)
migrate-schema:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose up
migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose up $(step)
migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose down $(step)
migrate-force:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI)  force $(version)