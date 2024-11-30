///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Создайте Makefile для музыкального сервиса и загрузите в этот Makefile следующее:
YOUR_POSTGRESQL_PASSWORD — ваш пароль Postgres.
YOUR_DB_NAME — ваша база данных Postgres.

gen-proto:
	protoc --go_out=. --go-grpc_out=. proto/music.proto

migrate-import:
	go get github.com/golang-migrate/migrate/v4/database/postgres

migrate-create:
	migrate create -ext sql -dir pkg/migrations -seq migrate

migrate-up:
	migrate -database postgres://postgres:YOUR_POSTGRESQL_PASSWORD@localhost:5432/YOUR_DB_NAME?sslmode=disable -path ./pkg/migrations up

migrate-down:
	migrate -database postgres://postgres:YOUR_POSTGRESQL_PASSWORD@localhost:5432/YOUR_DB_NAME?sslmode=disable -path ./pkg/migrations down

migrate-force:
	migrate -database postgres://postgres:YOUR_POSTGRESQL_PASSWORD@localhost:5432/YOUR_DB_NAME?sslmode=disable -path ./pkg/migrations force 1

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Создайте файл .env для музыкального сервиса и загрузите следующее:

dataSourceName="postgres://postgres:YOUR_POSTGRESQL_PASSWORD@localhost:5432/YOUR_DB_NAME?sslmode=disable"

SONG_PORT="50070"

clientID="clientID приложения в ВАШЕМ SPOTIFY ДЛЯ РАЗРАБОТЧИКОВ"
clientSecret="clientID приложения в ВАШЕМ SPOTIFY ДЛЯ РАЗРАБОТЧИКОВ"

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

Все сервисы запускаете отдельно, основной старт музыкального сервиса находится в файле main.go CMD/MAIN>GO, а API-шлюза - API/CMD/MAIN.GO