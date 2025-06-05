# Makefile для создания миграций


DB_DSN := "postgres://postgres:qwerty@localhost:5438/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)


migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	$(MIGRATE) up


migrate-down:
	$(MIGRATE) down


run:
	go run cmd/app/main.go