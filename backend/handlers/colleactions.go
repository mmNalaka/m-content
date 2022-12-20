package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/m-content/db"
	"github.com/mmnalaka/m-content/db/sqlc"
	"github.com/mmnalaka/m-content/models"
	"github.com/mmnalaka/m-content/utils"
)

// Content management API handlers
func CreateCollection(c *fiber.Ctx) error {
	args := models.CollocationsCreateRequest{}

	// Bind the request body to the args struct
	if err := c.BodyParser(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body," + err.Error(),
			"status":  "error",
			"code":    fiber.StatusBadRequest,
		})
	}

	// Validate the request body
	v := utils.NewValidator()
	if err := v.Struct(args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body," + err.Error(),
			"status":  "error",
			"code":    utils.ValidatorErrors(err),
		})
	}

	// Validate the slug
	if !utils.IsValidSlug(args.Slug) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid slug format",
			"status":  "error",
			"code":    fiber.StatusBadRequest,
		})
	}

	collection, err := db.Store.CreateCollection(c.Context(), sqlc.CreateCollectionParams{
		Name:      sql.NullString{String: args.Name, Valid: true},
		Slug:      args.Slug,
		Notes:     sql.NullString{String: args.Notes, Valid: true},
		Singleton: args.Singleton,
		CreatedBy: args.CreatedBy,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating collection," + err.Error(),
			"status":  "error",
			"code":    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(collection)
}
