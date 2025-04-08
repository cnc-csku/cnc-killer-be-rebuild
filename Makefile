include .env

DB_URI="postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:$(EXTERNAL_DB_PORT)/$(DB_DATABASE)?sslmode=$(DB_SSL_MODE)"
MIGRATIONS_PATH=internal/migrations
step=1
migrate-create:
	migrate create -ext=sql -dir=$(MIGRATIONS_PATH) -tz "UTC" $(name)
migrate-all:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose up
migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose up $(step)
migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI) -verbose down $(step)
migrate-force:
	migrate -path=$(MIGRATIONS_PATH) -database $(DB_URI)  force $(version)