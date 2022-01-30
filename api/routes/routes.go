package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtito/tinyApp/controller"
	"github.com/iamtito/tinyApp/middleware"
)

func Setup(app *fiber.App) {
	app.Get("/", controller.Home)
	app.Get("/other", controller.Other)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/forgot", controller.ForgotPasswordRequest)
	app.Post("/api/reset", controller.ResetPassword)

	app.Use(middleware.IsAuthenticated)

	app.Put("/api/users/info", controller.UpdateInfo)
	app.Put("/api/users/password", controller.UpdatePassword)

	app.Get("/api/user", controller.User)
	app.Get("/api/users", controller.AllUsers)
	app.Get("/api/users/:id", controller.GetUser)
	app.Post("/api/users", controller.CreateUser)
	app.Put("/api/users/:id", controller.UpdateUser)
	app.Delete("/api/users/:id", controller.DeleteUser)
	app.Post("/api/logout", controller.Logout)
	// app.Post("/api/users", controller.CreateUser)

	app.Get("/api/roles", controller.AllRoles)
	app.Post("/api/role", controller.CreateRole)
	app.Get("/api/role/:id", controller.GetRole)
	app.Put("/api/role/:id", controller.UpdateRole)
	app.Delete("/api/role/:id", controller.DeleteRole)

	app.Get("/api/permissions", controller.AllPermissions)

	app.Get("/api/products", controller.AllProducts)
	app.Post("/api/product", controller.CreateProduct)
	app.Get("/api/product/:id", controller.GetProduct)
	app.Put("/api/product/:id", controller.UpdateProduct)
	app.Delete("/api/product/:id", controller.DeleteProduct)

	app.Post("/api/upload", controller.Upload)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/orders", controller.AllOrders)
	app.Post("/api/export", controller.Export)
	app.Get("/api/chart", controller.Chart)
}
