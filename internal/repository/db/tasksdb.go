package database

import (
	"github.com/jmoiron/sqlx"
	"wbl0/internal/repository/model"
)

type AddDb struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *AddDb {
	return &AddDb{db: db}
}

func (db *AddDb) SaveOrder(data *model.DbItem) (string, error) {
	//some code

	return "ok", nil
}

func (db *AddDb) GetAllOrders() ([]model.DbItem, error) {
	//some code

	return []model.DbItem{}, nil
}
func (db *AddDb) GetOrderByID(id string) (model.DbItem, error) {
	//some code, select data from map

	return model.DbItem{}, nil
}
