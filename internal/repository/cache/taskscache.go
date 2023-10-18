package cache

import (
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
	log "github.com/sirupsen/logrus"
)

type AddCache struct {
	Cache
}

func (c *AddCache) AddToCache(data model.OrderData) {
	log.Infoln("Add data to cache")
	c.Cache[data.OrderUid] = data
}

func (c *AddCache) GetFromCache(id string) model.OrderData {
	log.Infoln("Get data from cache")
	return c.Cache[id]
}
