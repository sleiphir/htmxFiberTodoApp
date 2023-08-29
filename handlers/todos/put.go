package todoHandler

import (
    "fiberTodo/repositories"
    "strings"

    "github.com/gofiber/fiber/v2"
)

func PutEdit(c *fiber.Ctx) error {
    todoStore := repositories.GetTodoRepository()
    todos, _ := todoStore.Load()
    id := c.Params("id")
    if _, ok := todos[id]; !ok {
        return c.Status(404).SendString("Not Found")
    }
    title := strings.Clone(c.FormValue("title"))
    if title == "" {
        return c.Status(400).SendString("No title provided")
    }
    completed := c.FormValue("completed") == "on"
    todos[id].Title = title
    todos[id].Completed = completed
    todoStore.Save(todos)
    // infoLog.Println("Updated todo with id: " + id)
    return c.Render("partials/todo/item", fiber.Map{
        "Id": todos[id].Id,
        "Title": todos[id].Title,
        "Completed": todos[id].Completed,
    })
}
