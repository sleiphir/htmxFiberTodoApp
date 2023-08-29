package handlers

import (
	"fiberTodo/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

var todoStore *repositories.TodoStore = repositories.GetTodoRepository()

func IndexTodos(c *fiber.Ctx) error {
    todos, err := todoStore.Load()
    if err != nil {
        log.Fatal(err)
    }
    return c.Render("index", fiber.Map{
        "Title": "Todos",
        "Todos": todos,
    }, "layouts/main")
}
