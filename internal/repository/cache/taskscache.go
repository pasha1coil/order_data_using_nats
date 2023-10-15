package cache

import "wbl0/internal/repository/model"

type AddCache struct {
	Cache
}

func (c *Cache) AddToCache(data model.OrderData) {
	//some code
}

func (c *Cache) GetFromCache(id string) model.OrderData {
	//some code

	return model.OrderData{}
}
