package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/m-content/handlers"
)

func SetupAppRoutes(app *fiber.App) {
	// app routes
	app.Get("/health", PlaceholderHandler)
	app.Get("/version", PlaceholderHandler)
	app.Get("/ready", PlaceholderHandler)

	// api v1 routes
	v1 := app.Group("/api/v1")

	// auth routers
	auth := v1.Group("/auth")
	auth.Post("/login", handlers.AuthLoinHandler)
	auth.Post("/register", handlers.AuthRegisterHandler)
	auth.Post("/logout", PlaceholderHandler)
	auth.Post("/refresh", PlaceholderHandler)
	auth.Post("/userinfo", PlaceholderHandler)
	auth.Post("/verify", PlaceholderHandler)
	auth.Post("/change-password", PlaceholderHandler)
	auth.Post("/forgot-password", PlaceholderHandler)
	auth.Post("/reset-password/done", PlaceholderHandler)

	// Protected routes
	cda := v1.Group("/cda", ApiKeyProtected())
	cda.Get("/collections", PlaceholderHandler)
	cda.Get("/collections/:id", PlaceholderHandler)

	// Protected routes
	cma := v1.Group("/cma", JWTProtected())
	cma.Get("/collections", PlaceholderHandler)
	cma.Get("/collections/:id", PlaceholderHandler)
	cma.Post("/collections", handlers.CreateCollection)
	cma.Put("/collections/:id", PlaceholderHandler)
	cma.Delete("/collections/:id", PlaceholderHandler)
}

// dummy handlers
func PlaceholderHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World 👋!",
		"path:":   c.Path(),
		"method":  c.Method(),
	})
}

// ApiKeyProtected dummy middleware
func ApiKeyProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

// JWT protected dummy middleware
func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
