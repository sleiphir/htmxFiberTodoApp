package main

import (
    "encoding/json"
    "github.com/google/uuid"
    "log"
    "os"
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/storage/sqlite3/v2"
    "github.com/gofiber/template/html/v2"
)

type Todo struct {
    Id string
    Title string
    Completed bool
}

type Todos map[string]*Todo

type TodoStore struct {
    Store *sqlite3.Storage
}

const (
    todosKey = "todos"
)

var id = 0

func (t *TodoStore) Load() (Todos, error) {
    rawTodos, err := t.Store.Get(todosKey)
    if err != nil {
        return Todos{}, err
    }
    todos := &map[string]*Todo{}
    json.Unmarshal(rawTodos, todos)
    return *todos, nil
}

func (t *TodoStore) Save(val Todos) error {
    rawTodos, err := json.Marshal(val)
    if err != nil {
        return err
    }
    return t.Store.Set(todosKey, rawTodos, 0)
}

func main() {
    infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

    todoStore := TodoStore{
        Store: sqlite3.New(),
    }

    todos, err := todoStore.Load()
    if err != nil {
        log.Fatal(err)
    }

    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Todos",
            "Todos": todos,
        }, "layouts/main")
    })

    app.Get("/todos/:id/edit", func(c *fiber.Ctx) error {
        id := c.Params("id")
        if _, ok := todos[id]; !ok {
            return c.Status(404).SendString("Not Found")
        }
        return c.Render("partials/todo/edit_item", fiber.Map{
            "Id": todos[id].Id,
            "Title": todos[id].Title,
            "Completed": todos[id].Completed,
        })
    })

    app.Post("/todos/:id/toggle", func(c *fiber.Ctx) error {
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
    })

    app.Post("/todos/new", func(c *fiber.Ctx) error {
        id := uuid.New().String()
        title := strings.Clone(c.FormValue("title"))
        if title == "" {
            return c.Status(400).SendString("No title provided")
        }
        todos[id] = &Todo{
            Id: id,
            Title: title,
            Completed: false,
        }
        todoStore.Save(todos)
        infoLog.Println("Added new todo with id: " + id)
        return c.Render("partials/todo/new_item_added", fiber.Map{
            "Id": todos[id].Id,
            "Title": todos[id].Title,
            "Completed": todos[id].Completed,
        })
    })

    app.Put("/todos/:id", func(c *fiber.Ctx) error {
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
        infoLog.Println("Updated todo with id: " + id)
        return c.Render("partials/todo/item", fiber.Map{
            "Id": todos[id].Id,
            "Title": todos[id].Title,
            "Completed": todos[id].Completed,
        })
    })

    app.Delete("/todos/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        delete(todos, id)
        todoStore.Save(todos)
        infoLog.Println("Deleted todo with id: " + id)
        return c.SendString("")
    })

    log.Fatal(app.Listen(":3000"))
}

