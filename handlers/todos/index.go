package todoHandler

import (
	"github.com/gofiber/fiber/v2"
)

func GetGroup(app *fiber.App) fiber.Router {
    todoHandler := app.Group("/todos")

    todoHandler.Get(":id/edit", GetEdit)
    todoHandler.Post(":id/toggle", PostToggle)
    todoHandler.Post("new", PostNew)
    todoHandler.Put(":id", PutEdit)
    todoHandler.Delete(":id", Delete)

    return todoHandler
}
