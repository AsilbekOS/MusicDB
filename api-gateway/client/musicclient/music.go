package musicclient

import (
	"apigateway/proto/musicproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewUserServiceClient foydalanuvchi xizmatiga gRPC ulanishini yaratadi va UserServiceClient-ni qaytaradi.
func NewSongServiceClient(userAddress string) (musicproto.SongServiceClient, error) {
	conn, err := grpc.NewClient(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to music-service: %v\n", err)
		return nil, err
	}

	user := musicproto.NewSongServiceClient(conn)
	return user, nil
}
