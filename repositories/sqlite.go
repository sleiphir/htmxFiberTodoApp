package repositories

import (
	"fmt"
	"sync"

	"github.com/go-rel/rel"
	"github.com/go-rel/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
    repository rel.Repository
}

const DB_FILE = "dev.db"
var lock = &sync.Mutex{}
var repository *rel.Repository

func GetRepository() (*rel.Repository, error) {
    if repository == nil {
        lock.Lock()
        defer lock.Unlock()
        if repository == nil {
            var err error
            fmt.Println("Instantiating new repository")
            repository, err = initialize()
            if err != nil {
                return nil, err
            }
        }
    } else {
        fmt.Println("Repository already instantiated")
    }
    return repository, nil
}

func initialize() (*rel.Repository, error) {
	adapter, err := sqlite3.Open(DB_FILE)
	if err != nil {
        return nil, err
	}
	//defer adapter.Close()

    repo := rel.New(adapter)
    return &repo, nil
}
