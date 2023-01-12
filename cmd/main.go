package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/eserzhan/test-task/database"

)

func main() {
    database.ConnectDb()
    app := fiber.New()

    app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

    setupRoutes(app)

    app.Listen(":3000")
}