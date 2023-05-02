package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func User(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies(("jwt"))

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	payload := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", payload.Subject).First(&user)

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
