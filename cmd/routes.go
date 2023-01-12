package main

import (
	// "github.com/eserzhan/test-task/handlers"
	"github.com/gofiber/fiber/v2"
    "github.com/eserzhan/test-task/controllers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/user", controllers.ListUsers)

    app.Post("/user", controllers.CreateUser)

    app.Delete("/user/:id", controllers.DeleteUser)

    app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
}