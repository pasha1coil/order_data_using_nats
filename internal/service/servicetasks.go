package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/pasha1coil/order_data_using_nats/internal/repository"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
	log "github.com/sirupsen/logrus"
)

type AddService struct {
	repo repository.Repository
}

func NewAddService(repo repository.Repository) *AddService {
	return &AddService{
		repo: repo,
	}
}

// Сохранить данные заказа
func (a *AddService) SaveOrderData(data []byte) error {
	orderdata := new(model.OrderData)
	err := orderdata.Get(data)
	if err != nil {
		log.Errorf("Wrong format: %s", err.Error())
		return err
	}
	validate := validator.New()
	err = validate.Struct(orderdata)
	if err != nil {
		log.Errorf("Error validate: %s", err.Error())
		return err
	}
	itemData := new(model.DbItem)
	itemData.Order = *orderdata
	itemData.Id = orderdata.OrderUid
	a.repo.AddToCache(*orderdata)
	_, err = a.repo.SaveOrder(itemData)
	if err != nil {
		log.Errorf("Erroe save order:%s", err.Error())
		return err
	}
	return err
}

// Получить из кэша по UID
func (a *AddService) GetFromCacheByUID(id string) model.OrderData {
	return a.repo.GetFromCache(id)
}

// Получить все заказы
func (a *AddService) GetAllOrders() ([]model.DbItem, error) {
	di, err := a.repo.GetAllOrders()
	if err != nil {
		log.Errorf("Error get all orders:%s", err.Error())
		return nil, err
	}
	return di, err
}

// Восстановить кэш
func (a *AddService) RestoreCache() error {
	dbItems, err := a.GetAllOrders()
	if dbItems == nil {
		log.Errorf("Error get all orders from RestoreCache:%s", err.Error())
		return err
	}
	for _, dbItem := range dbItems {
		a.repo.AddToCache(dbItem.Order)
	}
	log.Infoln("Cache is restored")
	return err
}
