package cache

import (
	log "github.com/sirupsen/logrus"
	"wbl0/internal/repository/model"
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
