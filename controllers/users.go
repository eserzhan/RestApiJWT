package controllers

import (

	"github.com/eserzhan/test-task/database"
	"github.com/eserzhan/test-task/models"
	"github.com/gofiber/fiber/v2"
)

func ListUsers(c *fiber.Ctx) error {
    Users := []models.User{}
    database.DB.Db.Find(&Users)

    return c.Status(200).JSON(Users)
}

func CreateUser(c *fiber.Ctx) error {
    User := new(models.User)
    if err := c.BodyParser(User); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&User)

    return c.Status(200).JSON(User)
}

func DeleteUser(c * fiber.Ctx) error {
    id := c.Params("id")
    var User models.User
    res := database.DB.Db.Delete(&User, id)

    if res.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}