generate-user:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/musicproto/music.proto

swag-gen:
	~/go/bin/swag init -g ./api/cmd/main.go -o api/docs force 1	