package todoHandler

import (
	"fiberTodo/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetEdit(c *fiber.Ctx) error {
    todos, _ := repositories.GetTodoRepository().Load()
    id := c.Params("id")
    if _, ok := todos[id]; !ok {
        return c.Status(404).SendString("Not Found")
    }
    return c.Render("partials/todo/edit_item", fiber.Map{
        "Id": todos[id].Id,
        "Title": todos[id].Title,
        "Completed": todos[id].Completed,
    })
}
