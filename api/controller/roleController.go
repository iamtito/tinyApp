package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/users"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []users.Role

	database.DB.Find(&roles)
	return c.JSON(roles)
}

// type RoleCreateDTO struct {
// 	Name        string
// 	Permissions []int
// 	// Permissions []json.RawMessage
// }

func CreateRole(c *fiber.Ctx) error {
	// var role users.Role
	// var roleDto RoleCreateDTO
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		log.Println(err)
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]users.Permissions, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		// id := permissionId
		permissions[i] = users.Permissions{
			ID: uint(id),
		}
	}

	role := users.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var role users.Role
	role.ID = uint(id)
	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		log.Println(err)
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permissions := make([]users.Permissions, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		// id := permissionId
		permissions[i] = users.Permissions{
			ID: uint(id),
		}
	}
	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := users.Role{
		ID:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}
	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var role users.Role
	role.ID = uint(id)
	database.DB.Where("id = ?", role.ID).First(&role)
	if role.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Role id:" + strconv.FormatInt(int64(role.ID), 10) + " does not exist!",
		})
	}
	database.DB.Delete(&role)
	return nil
}
