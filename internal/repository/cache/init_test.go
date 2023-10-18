package cache

import (
	"testing"

	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
)

func TestInitCache(t *testing.T) {
	cache := InitCache()

	if cache == nil {
		t.Fatalf("Expected cache to be initialized, got nil")
	}

	if len(cache.Cache) != 0 {
		t.Errorf("Expected cache to be empty on initialization, got size: %d", len(cache.Cache))
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := InitCache()

	orderData := model.OrderData{
		OrderUid: "1",
	}

	cache.Cache["key"] = orderData

	if len(cache.Cache) != 1 {
		t.Errorf("Expected cache size to be 1 after adding an item, got size: %d", len(cache.Cache))
	}

	retrievedOrder, ok := cache.Cache["key"]
	if !ok {
		t.Errorf("Failed to get order from cache")
	}
	if retrievedOrder.OrderUid != "1" {
		t.Errorf("Retrieved order does not match added order")
	}
}
