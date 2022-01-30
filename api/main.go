package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iamtito/tinyApp/database"
	"github.com/iamtito/tinyApp/routes"
	_ "github.com/iamtito/tinyApp/users"
)

func main() {

	database.Connect()
	// user := users.User{
	// 	Name: "Tito",
	// }
	// fmt.Println(user)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":9080")

}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("you canr de")
	}
	return a / b, nil
}
