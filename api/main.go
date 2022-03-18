package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

/* register all your routes */
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL) /* when we hit the hit /api/v1 with POST request, we send the URL to be shortened to this func(does all the magic) */
}

func main() {
	err := godotenv.Load() /* to load my environment variables */

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New()) /* to log HTTP request/response details */

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT"))) /* create/start the server here */
}
