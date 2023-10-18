package service

import (
	"github.com/pasha1coil/order_data_using_nats/internal/repository"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type Service struct {
	Tasks
}

type Tasks interface {
	GetFromCacheByUID(id string) model.OrderData
	GetAllOrders() ([]model.DbItem, error)
	SaveOrderData(data []byte) error
	RestoreCache() error
}

func NewService(repo *repository.MainRepository) *Service {
	return &Service{
		Tasks: NewAddService(repo.Repository),
	}
}
