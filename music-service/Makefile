gen-proto:
	protoc --go_out=. --go-grpc_out=. proto/music.proto

migrate-import:
	go get github.com/golang-migrate/migrate/v4/database/postgres

migrate-create:
	migrate create -ext sql -dir pkg/migrations -seq migrate

migrate-up:
	migrate -database postgres://postgres:1915@localhost:5432/musicdb?sslmode=disable -path ./pkg/migrations up

migrate-down:
	migrate -database postgres://postgres:1915@localhost:5432/musicdb?sslmode=disable -path ./pkg/migrations down

migrate-force:
	migrate -database postgres://postgres:1915@localhost:5432/musicdb?sslmode=disable -path ./pkg/migrations force 1
