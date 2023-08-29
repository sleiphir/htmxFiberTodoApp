package todoHandler

import (
	"fiberTodo/models"
	"fiberTodo/repositories"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostToggle(c *fiber.Ctx) error {
    todoStore := repositories.GetTodoRepository()
    todos, _ := todoStore.Load()
    id := c.Params("id")
    if _, ok := todos[id]; !ok {
        return c.Status(404).SendString("Not Found")
    }
    todos[id].Completed = !todos[id].Completed
    todoStore.Save(todos)
    return c.Render("partials/todo/item", fiber.Map{
        "Id": todos[id].Id,
        "Title": todos[id].Title,
        "Completed": todos[id].Completed,
    })
}

func PostNew(c *fiber.Ctx) error {
        todoStore := repositories.GetTodoRepository()
        todos, _ := todoStore.Load()
        id := uuid.New().String()
        title := strings.Clone(c.FormValue("title"))
        if title == "" {
            return c.Status(400).SendString("No title provided")
        }
        todos[id] = &models.Todo{
            Id: id,
            Title: title,
            Completed: false,
        }
        todoStore.Save(todos)
        // infoLog.Println("Added new todo with id: " + id)
        return c.Render("partials/todo/new_item_added", fiber.Map{
            "Id": todos[id].Id,
            "Title": todos[id].Title,
            "Completed": todos[id].Completed,
        })
    }
