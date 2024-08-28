package handlers

import (
	"github.com/gofiber/fiber/v2"
	"jwt_auth_fiber/models"
	"jwt_auth_fiber/utils"
)

const (
	badRequestMessage = "Something's wrong with your input"
	unauthorized      = "unauthorized"
	successMessage    = "Welcome to the protected area"
	status200         = fiber.StatusOK
	status400         = fiber.StatusBadRequest
	status401         = fiber.StatusUnauthorized
	status500         = fiber.StatusInternalServerError
)

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(status400).JSON(models.Response{Message: badRequestMessage, Status: status400})
	}

	if !models.IsThereUser(user) {
		return c.Status(status401).JSON(models.Response{Message: unauthorized, Status: status401})
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return c.Status(status500).JSON(models.Response{Message: err.Error(), Status: status500})
	}

	loginResponse := models.LoginResponse{Token: token, TokenType: "bearer"}
	return c.Status(status200).JSON(loginResponse)
}

func Protected(c *fiber.Ctx) error {
	response := models.Response{Message: successMessage, Status: status200}
	return c.Status(status200).JSON(response)
}
