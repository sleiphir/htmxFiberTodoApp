package migrations

import (
	"context"
	"fiberTodo/models"

	"github.com/go-rel/rel"
)

// MigrateCreateTodos definition
func MigrateCreateTodos(schema *rel.Schema) {
    schema.CreateTable("todos", func(t *rel.Table) {
        t.ID("id")
        t.String("title")
        t.Bool("completed")
        t.DateTime("created_at")
        t.DateTime("updated_at")
    })

    schema.Do(func(ctx context.Context, repo rel.Repository) error {
        // add seeds
        return repo.Insert(ctx, &models.Todo{Title: "Learn Go"})
    })
}

// RollbackCreateTodos definition
func RollbackCreateTodos(schema *rel.Schema) {
    schema.DropTable("todos")
}
