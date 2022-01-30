package controller

import (
	"log"
	"math/rand"

	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/users"
	"golang.org/x/crypto/bcrypt"
)

func ForgotPasswordRequest(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return nil
	}
	token := RandStringRunes(12)
	log.Println(token)
	passwordReset := users.PasswordReset{
		Email: data["email"],
		Token: token,
	}
	user := users.User{
		Email: data["email"],
	}
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email does not exist exist!",
		})
	}
	database.DB.Create(&passwordReset)

	from := "admin@example.com"

	to := []string{
		data["email"],
	}

	url := "http://127.0.0.1:3000/reset/" + token

	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password!")
	// message := := []byte("To: "+data["email"]+"\r\n" +
	// 					"Subject: discount Gophers!\r\n" +
	// 					"\r\n" +
	// 					"This is the email body.\r\n")
	err := smtp.SendMail("mail:1025", nil, from, to, message)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ResetPassword(c *fiber.Ctx) error {
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
	var passwordReset = users.PasswordReset{}
	if err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token!",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	database.DB.Model(&users.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "password reset success",
	})
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
