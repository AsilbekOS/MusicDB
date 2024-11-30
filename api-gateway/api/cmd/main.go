package main

import (
	"apigateway/api/router"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "apigateway/api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/swagger"
)

// New ...
// @title MusicDB
// @version 1.0
// @description This is a API for demonstration purposes.
// @host localhost:50077
// @BasePath /
// @contact.name Asilbek Kholmatov
// @contact.email asilbekxolmatov2002@gmail.com
func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Swagger UI setup for API documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// CORS configuration: allows requests from specific origins
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://192.168.0.116:50077, https://localhost:50077", // Only these origins are allowed
		AllowMethods:     "GET,POST,DELETE,PUT,PATCH,OPTIONS",                    // Allowed methods
		AllowHeaders:     "Content-Type, api_key, Authorization",                 // Allowed headers
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length", // Expose headers to the client
	}))

	// Middleware to log incoming requests
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s", c.Method(), c.OriginalURL())
		return c.Next()
	})

	// Rate Limiting middleware: allows only 5 requests per 30 seconds for each IP
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        5,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html") // Send custom "Too fast" page if limit is exceeded
		},
	}))

	// Initialize routing for user-related song API
	router.UserRouter(app)

	// Error handling middleware: Logs and handles errors
	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString("An error occurred. Please try again later.")
		}
		return nil
	})

	// Graceful shutdown when termination signal is received
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		<-quit
		log.Println("Gracefully shutting down...")
		app.ShutdownWithContext(context.Background()) // Shutdown the server
	}()
	defer close(quit)

	// Start the server with HTTPS (TLS enabled)
	if err := app.ListenTLS("localhost:50077", "api/certs/cert.pem", "api/certs/key.pem"); err != nil {
		log.Panic(err) // Log error if the server fails to start
	}
}
