package cache

import (
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
	"testing"
)

func TestAddToCache(t *testing.T) {
	cache := InitCache()

	orderData := model.OrderData{
		OrderUid: "Test",
	}

	cache.AddToCache(orderData)

	if len(cache.Cache) != 1 {
		t.Errorf("Expected cache size to be 1 after adding an item, got size: %d", len(cache.Cache))
	}

	if cache.Cache[orderData.OrderUid].OrderUid != "Test" {
		t.Errorf("Failed to retrieve correct data from cache")
	}
}

func TestGetFromCache(t *testing.T) {
	cache := InitCache()

	orderData := model.OrderData{
		OrderUid: "Test",
	}

	cache.AddToCache(orderData)

	retrievedData := cache.GetFromCache(orderData.OrderUid)

	if retrievedData.OrderUid != "Test" {
		t.Errorf("Failed to retrieve correct data from cache")
	}
}
