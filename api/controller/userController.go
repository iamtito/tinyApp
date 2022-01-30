package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/middleware"
	"github.com/iamtito/tinyApp/users"
)

const DefaultPassword = "1234"

func AllUsers(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(users.Paginate(database.DB, &users.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user users.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Where("email = ?", user.Email).First(&user)
	if user.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User with similar email already exist!",
		})
	}
	user.SetPassword(DefaultPassword)

	// user.Password = password
	log.Println("holla")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	var user users.User
	user.ID = uint(id)
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	var user users.User
	user.ID = uint(id)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	var user users.User
	user.ID = uint(id)
	database.DB.Where("email = ?", user.Email).First(&user)
	if user.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User does not exist!",
		})
	}
	database.DB.Delete(&user)
	return nil
}
