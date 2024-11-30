package router

import (
	"apigateway/api/handler"
	"apigateway/client/musicclient"
	"apigateway/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func UserRouter(app *fiber.App) {
	// Connecting to the user via gRPC
	uClinet, err := musicclient.NewSongServiceClient("localhost:50070")
	if err != nil {
		log.Println(err)
	}
	songHandler := handler.NewSongClient(uClinet)
	// log.Println("Arrived: Router: 2")

	// Routers and middleware
	api := app.Group("/api/songs")
	api.Post("/add", middleware.FiberCORSMiddleware(songHandler.AddSong))
	api.Get("/id", middleware.FiberCORSMiddleware(songHandler.GetSong))
	api.Get("/all", middleware.FiberCORSMiddleware(songHandler.GetSongs))
	api.Put("/update", middleware.FiberCORSMiddleware(songHandler.UpdateSong))
	api.Delete("/delete", middleware.FiberCORSMiddleware(songHandler.DeleteSong))
	api.Get("/lyrics/id", middleware.FiberCORSMiddleware(songHandler.GetSongLyrics))

}
