package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"wbl0/internal/repository/cache"
	database "wbl0/internal/repository/db"
	"wbl0/internal/repository/model"
)

type Store struct {
	cache cache.AddCache
	db    database.AddDb
}

func InitStore(cache cache.AddCache, db database.AddDb) *Store {
	return &Store{
		cache: cache,
		db:    db,
	}
}

// Сохранить данные заказа
func (ss *Store) SaveOrderData(data []byte) error {
	od := new(model.OrderData)
	err := od.Get(data)
	if err != nil {
		log.Println("Wrong format")
		return err
	}
	validate := validator.New()
	err = validate.Struct(od)
	if err != nil {
		log.Println(err)
		return err
	}
	itemData := new(model.DbItem)
	itemData.Order = *od
	itemData.Id = od.OrderUid
	ss.cache.AddToCache(*od)
	_, err = ss.db.SaveOrder(itemData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Получить из кэша по UID
func (ss *Store) GetFromCacheByUID(id string) model.OrderData {
	return ss.cache.GetFromCache(id)
}

// Получить все заказы
func (ss *Store) GetAllOrders() ([]model.DbItem, error) {
	di, err := ss.db.GetAllOrders()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return di, err
}

// восстановить кэш
func (ss *Store) RestoreCache() error {
	dItems, err := ss.GetAllOrders()
	if dItems == nil {
		log.Println(err)
		return err
	}
	for _, dItem := range dItems {
		ss.cache.AddToCache(dItem.Order)
	}
	log.Println("--CACHE IS RESTORED--")
	return err
}
