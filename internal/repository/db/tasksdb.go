package database

import (
	"context"
	"encoding/json"
	"fmt"
	"wbl0/internal/repository/model"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AddDb struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewDB(db *sqlx.DB) *AddDb {
	return &AddDb{db: db, ctx: context.Background()}
}

func (d *AddDb) SaveOrder(data *model.DbItem) (string, error) {
	log.Infoln("Save order to database")

	jsonOrderData, err := json.Marshal(data.Order)
	if err != nil {
		return "", fmt.Errorf("error marshalling order data: %w", err)
	}

	_, err = d.db.ExecContext(d.ctx, "INSERT INTO Orders (Id, Ord) values ($1,$2)", data.Id, string(jsonOrderData))
	if err != nil {
		log.Errorf("Error saving order to database:%v", err)
		return "", fmt.Errorf("error saving order to database: %w", err)
	}
	log.Infoln("OK")
	return "OK", nil
}

func (d *AddDb) GetAllOrders() ([]model.DbItem, error) {
	log.Infoln("Get all orders from database")
	data := []model.DbItem{}
	rows, err := d.db.QueryContext(d.ctx, "SELECT * FROM Orders")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			log.Errorf("Error closing rows: %v", err)
		}
	}()

	for rows.Next() {
		var str model.DbItem
		var orderData string
		if err := rows.Scan(&str.Id, &orderData); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		order := &model.OrderData{}
		if err := json.Unmarshal([]byte(orderData), order); err != nil {
			return nil, fmt.Errorf("error unmarshalling order data: %w", err)
		}
		str.Order = *order
		data = append(data, str)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("Error fetching rows: %v", err)
		return nil, fmt.Errorf("error fetching rows: %w", err)
	}
	return data, nil
}
