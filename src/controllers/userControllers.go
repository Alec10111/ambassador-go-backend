package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {

	data := ctx.AllParams()

	var user models.User
	database.DB.First(&user, "id = ?", data["id"])

	if user.Id == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return ctx.JSON(user)
}

// TODO: finish
func UpdateUserInfo(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.First(&user, data["id"])

	if user.Id == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// TODO: use Updates instead of loop of db queries
	for field, value := range data {
		database.DB.Model(&user).Update(field, value)
	}

	return ctx.JSON(user)

}
