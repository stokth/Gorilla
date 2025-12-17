# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)
PROTOS  := ./project-protos/proto/*.proto
OUT_DIR := .

# Генерации .proto файлов
generate:
	protoc \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTOS)

clean:
	find . -name "*.pb.go" -delete
# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up


# Откат миграций
migrate-down:
	$(MIGRATE) down

# Генерация кода
gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/tasks/api.gen.go

# Генерация кода users
gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/users/api.gen.go

# Проверка кода
lint:
	golangci-lint run --color=auto

# Тесты
test:
	go test ./... -v

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/main.go # Теперь при вызове make run мы запустим наш сервер