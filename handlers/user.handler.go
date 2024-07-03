package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/mail"
	"todo-app/configs"
	"todo-app/helpers"
	"todo-app/models"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func RegisterUser(c *fiber.Ctx) error {
	db := configs.DB
	user := new(models.User)

	// Store the body in the user and return error if encountered
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Check if email is already taken, if true return error
	if err := db.Where("email= ?", &user.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Email already taken", "data": err})
	}

	user.Password, _ = helpers.HashPassword(user.Password)

	err := db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func Login(c *fiber.Ctx) error {
	db := configs.DB
	var user models.User

	var input models.PayloadUserLogin

	// binding user input to a struct
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// set a variable depending on the condition
	var query string
	if valid(input.Email) {
		query = "email= ?"
	}

	if err := db.Where(query, input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error", "message": "User does not exists",
		})
	}

	if !helpers.ValidatePassword(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error", "messvalidage": "Password incorrect",
		})
	}

	token := helpers.GenerateToken(user)

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "token": token})
}
