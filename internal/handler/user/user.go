package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"go-web-service-demo/database"
	"go-web-service-demo/internal/model"
	"time"
)

// GetAll ShowAccount godoc
// @Summary      get all
// @Description  get all
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.UserDto
// @Router       /v1/users [get]
func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"data": nil})
	}

	return c.JSON(fiber.Map{"data": users})
}

func Create(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"data": err})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}

	return c.JSON(fiber.Map{"data": user})
}

func Get(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	id := c.Params("id")
	db.Find(&user, "id= ?", id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"data": nil})
	}
	return c.JSON(fiber.Map{"data": user})
}

func Update(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param noteId
	id := c.Params("id")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var userDto model.UserDto
	err := c.BodyParser(&userDto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"data": err})
	}

	// Edit the user
	user.Name = userDto.Name
	user.Surname = userDto.Surname
	user.Email = userDto.Email
	user.UpdatedAt = time.Now()

	// Save the Changes
	db.Save(&user)

	// Return the updated user
	return c.JSON(fiber.Map{"data": user})
}

func Delete(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"data": nil})
	}

	// Save the Changes
	db.Delete(&user)

	// Return the updated user
	c.Status(200)

	return nil
}
