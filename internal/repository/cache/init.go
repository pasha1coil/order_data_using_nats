package cache

import "wbl0/internal/repository/model"

type Cache map[string]model.OrderData

func InitCache() *AddCache {
	cache := make(Cache)
	return &AddCache{
		Cache: cache,
	}
}
