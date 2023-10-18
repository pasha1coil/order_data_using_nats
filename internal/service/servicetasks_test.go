package service

import (
	"github.com/golang/mock/gomock"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
	mock_service "github.com/pasha1coil/order_data_using_nats/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveOrderData(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTasks := mock_service.NewMockTasks(ctrl)

	bytes := []byte("string")
	mockTasks.EXPECT().SaveOrderData(bytes).Return(nil).Times(1)

	service := &Service{
		Tasks: mockTasks,
	}

	err := service.SaveOrderData(bytes)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestGetAllOrders(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTasks := mock_service.NewMockTasks(ctrl)

	items := []model.DbItem{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
	}

	mockTasks.EXPECT().GetAllOrders().Return(items, nil).Times(1)

	service := &Service{
		Tasks: mockTasks,
	}

	result, err := service.GetAllOrders()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != len(items) {
		t.Fatalf("Expected %d items but got %d", len(items), len(result))
	}
}

func TestGetFromCacheByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTasks := mock_service.NewMockTasks(ctrl)

	expectedResult := model.OrderData{OrderUid: "1"}
	mockTasks.EXPECT().GetFromCacheByUID("1").Return(expectedResult)

	testService := Service{
		Tasks: mockTasks,
	}

	result := testService.GetFromCacheByUID("1")

	assert.Equal(t, expectedResult, result)
}
