package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/users"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []users.Permissions

	database.DB.Find(&permissions)
	return c.JSON(permissions)
}
