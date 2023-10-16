package cache

import (
	log "github.com/sirupsen/logrus"
	"wbl0/internal/repository/model"
)

type Cache map[string]model.OrderData

func InitCache() *AddCache {
	log.Infoln("Init cache")
	cache := make(Cache)
	return &AddCache{
		Cache: cache,
	}
}
