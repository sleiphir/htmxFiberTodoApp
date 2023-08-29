package repositories

import (
	"encoding/json"
	"fiberTodo/models"

	"github.com/gofiber/storage/sqlite3/v2"
)

type TodoStore struct {
    Store *sqlite3.Storage
}

type Todos map[string]*models.Todo

const (
    todosKey = "todos"
)

var TodoRepository *TodoStore
var todos Todos // should be a pointer as maps are always passed by reference

func GetTodoRepository() *TodoStore {
    if TodoRepository == nil {
        TodoRepository = &TodoStore{
            Store: sqlite3.New(),
        }
    }
    return TodoRepository
}

func (t *TodoStore) Load() (Todos, error) {
    if todos != nil {
        return todos, nil
    }
    rawTodos, err := t.Store.Get(todosKey)
    if err != nil {
        return Todos{}, err
    }
    todosRef := &Todos{}
    json.Unmarshal(rawTodos, todosRef)
    todos = *todosRef
    return todos, nil
}

func (t *TodoStore) Save(val Todos) error {
    rawTodos, err := json.Marshal(val)
    if err != nil {
        return err
    }
    return t.Store.Set(todosKey, rawTodos, 0)
}
