package cache

import (
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
	log "github.com/sirupsen/logrus"
)

type Cache map[string]model.OrderData

func InitCache() *AddCache {
	log.Infoln("Init cache")
	cache := make(Cache)
	return &AddCache{
		Cache: cache,
	}
}
