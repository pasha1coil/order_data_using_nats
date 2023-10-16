package repository

import (
	"github.com/jmoiron/sqlx"
	"wbl0/internal/repository/cache"
	database "wbl0/internal/repository/db"
	"wbl0/internal/repository/model"
)

type MainRepository struct {
	Repository
}

type Repository struct {
	TasksDb
	TasksCache
}

// db
type TasksDb interface {
	SaveOrder(data *model.DbItem) (string, error)
	GetAllOrders() ([]model.DbItem, error)
}

// cache
type TasksCache interface {
	AddToCache(data model.OrderData)
	GetFromCache(id string) model.OrderData
}

func InitRepository(db *sqlx.DB) *MainRepository {
	return &MainRepository{
		Repository{
			TasksDb:    database.NewDB(db),
			TasksCache: cache.InitCache(),
		},
	}
}
