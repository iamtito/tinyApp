package controller

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/users"
	"github.com/iamtito/tinyApp/util"
	"gorm.io/gorm"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

func Other(c *fiber.Ctx) error {
	return c.SendString("Other functions")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// var user users.User

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}
	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := users.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}
	user.SetPassword(data["password"])

	result := database.DB.Where("email = ?", data["email"]).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		// fmt.Println(user)
		database.DB.Create(&user)

		return c.JSON(user)
		// return c.SendString("hi")
	} else {
		// if user.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User with similar email already exist!",
		})
	}
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user users.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
	// if err := user.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect email or password",
		})
	}
	// claims := jwt.StandardClaims{
	// 	// Id: strconv.Itoa(int(user.ID)),
	// 	Issuer:    strconv.Itoa(int(user.ID)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 days
	// }
	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    strconv.Itoa(int(user.ID)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 days
	// })

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID))) // claims.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logged in",
		"jwt":     token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user users.User
	database.DB.Preload("Role").Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logged out",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	user := users.User{
		ID:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	database.DB.Model(&user).Updates(user)
	// database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	// var user users.User
	user := users.User{}
	user.SetPassword(data["password"])
	database.DB.Model(&user).Where("id = ?", id).Updates(user)
	// database.DB.Where("id = ?", id).First(&user)
	return c.JSON(fiber.Map{
		"message": "password updated.",
	})
}
