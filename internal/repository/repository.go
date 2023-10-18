package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/cache"
	database "github.com/pasha1coil/order_data_using_nats/internal/repository/db"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

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
