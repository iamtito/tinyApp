package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/users"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	log.Println(c, page)
	return c.JSON(users.Paginate(database.DB, &users.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	var product users.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}
	// database.DB.Where("email = ?", user.Email).First(&user)
	// if user.ID != 0 {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": "User with similar email already exist!",
	// 	})
	// }

	// user.Password = password
	log.Println(product.Title, "product added")

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product users.Product
	product.Id = uint(id)
	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product users.Product
	product.Id = uint(id)
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Model(&product).Updates(product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product users.Product
	product.Id = uint(id)
	// database.DB.Where("email = ?", user.Email).First(&user)
	// if user.ID != 0 {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": "User does not exist!",
	// 	})
	// }
	database.DB.Delete(&product)
	return nil
}
