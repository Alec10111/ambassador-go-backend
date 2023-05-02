package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")

	// POST api/admin/register
	admin.Post("register", controllers.Register)

	// POST api/admin/login
	admin.Post("login", controllers.Login)

	// GET api/admin/user
	admin.Get("user", controllers.User)

	// TODO: POST api/admin/logout

	// TODO: PUT api/admin/users/info
	admin.Patch("users/:id", controllers.UpdateUserInfo)

	// TODO: PUT api/admin/password
}
