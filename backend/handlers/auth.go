package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/m-content/db"
	"github.com/mmnalaka/m-content/db/sqlc"
	"github.com/mmnalaka/m-content/models"
	"github.com/mmnalaka/m-content/utils"
)

// AuthRegisterHandler handles the auth.register request
func AuthRegisterHandler(c *fiber.Ctx) error {
	args := &models.AuthRegisterRequest{}

	// Check the request body
	if err := c.BodyParser(args); err != nil {
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

	passwordHash, err := utils.GeneratePasswordHash(args.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating password hash," + err.Error(),
			"status":  "error",
			"code":    fiber.StatusInternalServerError,
		})
	}

	// Create user in the database
	user, err := db.Store.CreateUser(c.Context(), sqlc.CreateUserParams{
		Email:    args.Email,
		Password: passwordHash,
		Name:     args.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user," + err.Error(),
			"status":  "error",
			"code":    fiber.StatusInternalServerError,
		})
	}

	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(models.AuthRegisterResponse{
		Message: "User created successfully",
		Status:  "success",
		Code:    fiber.StatusCreated,
		User: models.AuthRegisterUser{
			ID:        user.ID,
			Email:     user.Email,
			Password:  user.Password,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		},
	})
}

func AuthLoinHandler(c *fiber.Ctx) error {
	args := &models.AuthRegisterRequest{}

	// Check the request body
	if err := c.BodyParser(args); err != nil {
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

	// passwordHash, err := utils.GeneratePasswordHash(args.Password)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Error generating password hash," + err.Error(),
	// 		"status":  "error",
	// 		"code":    fiber.StatusInternalServerError,
	// 	})
	// }

	// // Create user in the database
	// user, err := db.Store.CreateUser(c.Context(), sqlc.CreateUserParams{
	// 	Email:    args.Email,
	// 	Password: passwordHash,
	// 	Name:     args.Name,
	// })
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Error creating user," + err.Error(),
	// 		"status":  "error",
	// 		"code":    fiber.StatusInternalServerError,
	// 	})
	// }

	// user.Password = ""

	// return c.Status(fiber.StatusCreated).JSON(models.AuthRegisterResponse{
	// 	Message: "User created successfully",
	// 	Status:  "success",
	// 	Code:    fiber.StatusCreated,
	// 	User: models.AuthRegisterUser{
	// 		ID:        user.ID,
	// 		Email:     user.Email,
	// 		Password:  user.Password,
	// 		Name:      user.Name,
	// 		CreatedAt: user.CreatedAt,
	// 	},
	// })

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"status":  "success",
		"code":    fiber.StatusCreated,
		"args":    args,
	})
}
