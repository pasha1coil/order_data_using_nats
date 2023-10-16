package service

import (
	"wbl0/internal/repository"
	"wbl0/internal/repository/model"
)

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
