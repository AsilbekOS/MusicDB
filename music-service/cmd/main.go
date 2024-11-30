package main

import (
	"log"
	"music/pkg/db"
	"music/proto"
	"music/service"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("SONG_PORT")
	if port == "" {
		log.Fatal("SONG_PORT is not set in environment variables")
	}

	// Ma'lumotlar bazasi
	// ----------------------------------------------------------------------------
	PostgreSQL, err := db.PostgreSQLConn()
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	// ----------------------------------------------------------------------------

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := service.NewSongService(PostgreSQL)
	s := grpc.NewServer()
	proto.RegisterSongServiceServer(s, server)
	reflection.Register(s)

	// Signal ushlab olish
	// ----------------------------------------------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		// 2 soniya kutamiz
		log.Println("Loading shutting down...")
		time.Sleep(2 * time.Second)
		s.GracefulStop()
		log.Println("Graceful shutting down server...")
	}()
	// ----------------------------------------------------------------------------

	log.Println("Server is listening on port: " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
