package cmd

import (
	"context"
	"fiberTodo/db/migrations"
	"fiberTodo/repositories"

	"github.com/go-rel/rel/migrator"
)

func main() {
    var (
        ctx  = context.TODO()
        repo, _ = repositories.GetRepository()
        m    = migrator.New(*repo)
    )

    // Register migrations
    m.Register(20230829173400, migrations.MigrateCreateTodos, migrations.RollbackCreateTodos)
    m.Migrate(ctx)
}
