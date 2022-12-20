package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/mmnalaka/m-content/config"
	"github.com/mmnalaka/m-content/db"
	"github.com/mmnalaka/m-content/router"
)

func main() {
	config := config.LoadConfig()

	app := fiber.New(config.Fiber)

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(cors.New(config.Cors))

	// Connect to database
	conn, err := db.OpenPostgresConnection(config.Postgres.URL)
	if err != nil {
		log.Fatal(err)
	}
	db.CreteNewStore(conn)

	// routes
	router.SetupAppRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
