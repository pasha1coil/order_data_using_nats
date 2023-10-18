package repository

import (
	"errors"
	mock_repository "github.com/pasha1coil/order_data_using_nats/internal/repository/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/model"
)

func TestSaveOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTaskDb := mock_repository.NewMockTasksDb(mockCtrl)
	mockTaskCache := mock_repository.NewMockTasksCache(mockCtrl)

	mockRepository := &MainRepository{Repository{
		TasksDb:    mockTaskDb,
		TasksCache: mockTaskCache}}

	testOrder := &model.DbItem{Id: "1", Order: model.OrderData{OrderUid: "Test"}}
	errTest := errors.New("error test")

	mockTaskDb.EXPECT().SaveOrder(testOrder).Return("1", errTest)
	_, err := mockRepository.SaveOrder(testOrder)

	if err != errTest {
		t.Errorf("returned err: got %v, want %v", err, errTest)
	}
}

func TestGetAllOrders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockTaskDb := mock_repository.NewMockTasksDb(mockCtrl)
	mockTaskCache := mock_repository.NewMockTasksCache(mockCtrl)

	mockRepository := &MainRepository{Repository{
		TasksDb:    mockTaskDb,
		TasksCache: mockTaskCache}}

	testOrders := []model.DbItem{{Id: "1", Order: model.OrderData{OrderUid: "Test1"}}, {Id: "2", Order: model.OrderData{OrderUid: "Test2"}}}
	errTest := errors.New("error test")

	mockTaskDb.EXPECT().GetAllOrders().Return(testOrders, errTest)
	orders, err := mockRepository.GetAllOrders()

	if err != errTest || !reflect.DeepEqual(orders, testOrders) {
		t.Errorf("returned: got %v, %v want %v, %v", orders, err, testOrders, errTest)
	}
}
