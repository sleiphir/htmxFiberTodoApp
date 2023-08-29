package todoHandler

import (
    "fiberTodo/repositories"

    "github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
    todoStore := repositories.GetTodoRepository()
    todos, _ := todoStore.Load()
    id := c.Params("id")
    delete(todos, id)
    todoStore.Save(todos)
    // infoLog.Println("Deleted todo with id: " + id)
    return c.SendString("")
}
