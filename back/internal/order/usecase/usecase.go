package usecase

import (
	"context"
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/order/repository"
	"net/http"
	"time"
)

type IOrderUsecase interface {
	OrdersPost(context.Context, models.Order) (models.Order, int, error)
	OrderUserIdGet(context.Context) (models.OrderList, int, error)
	OrderUserIdDelete(context.Context, uint64) (models.Order, int, error)
	DishIdPost(context.Context, uint64, uint64) (models.DishOrder, int, error)
	DishIdDelete(context.Context, uint64, uint64) (models.DishOrder, int, error)
}

type orderUsecase struct {
	OrderPostgresRepository repository.IOrderRepository
	Timeout                 time.Duration
}

func NewOrderUsecase(or repository.IOrderRepository, timeout time.Duration) IOrderUsecase {
	return &orderUsecase{
		OrderPostgresRepository: or,
		Timeout:                 timeout,
	}
}

func (ou *orderUsecase) OrdersPost(ctx context.Context, order models.Order) (models.Order, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.Order{}, http.StatusNotFound, nil
	}
	order.UserId = curUser.Id

	createdOrder, err := ou.OrderPostgresRepository.CreateOrder(ctx, order)
	if err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}
	return createdOrder, http.StatusOK, nil
}

func (ou *orderUsecase) OrderUserIdGet(ctx context.Context) (models.OrderList, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.OrderList{}, http.StatusNotFound, nil
	}

	orders, err := ou.OrderPostgresRepository.GetOrders(ctx, curUser.Id)
	if err != nil {
		return models.OrderList{}, http.StatusInternalServerError, err
	}
	for i := range orders {
		for j := range orders[i].Dishes {
			orders[i].Cost += orders[i].DishesCounts[j] * orders[i].DishesPrices[j]
		}
	}

	return orders, http.StatusOK, nil
}

func (ou *orderUsecase) OrderUserIdDelete(ctx context.Context, orderId uint64) (models.Order, int, error) {
	curUser, ok := ctx.Value(config.ContextUser).(models.User)
	if !ok {
		return models.Order{}, http.StatusNotFound, nil
	}

	deletedOrder, err := ou.OrderPostgresRepository.DeleteOrder(ctx, orderId, curUser.Id)
	if err != nil {
		return models.Order{}, http.StatusInternalServerError, err
	}
	return deletedOrder, http.StatusOK, nil
}

func (ou *orderUsecase) DishIdPost(ctx context.Context, dishId uint64, orderId uint64) (models.DishOrder, int, error) {
	_, err := ou.OrderPostgresRepository.GetOrderById(ctx, orderId)
	if err != nil {
		return models.DishOrder{}, http.StatusNotFound, err
	}

	number, err := ou.OrderPostgresRepository.AddDishToOrder(ctx, orderId, dishId)
	if err != nil {
		return models.DishOrder{}, http.StatusInternalServerError, err
	}

	dishes := models.DishOrder{
		Id:     dishId,
		Number: number,
	}

	return dishes, http.StatusOK, nil
}

func (ou *orderUsecase) DishIdDelete(ctx context.Context, dishId uint64, orderId uint64) (models.DishOrder, int, error) {
	_, err := ou.OrderPostgresRepository.GetOrderById(ctx, orderId)
	if err != nil {
		return models.DishOrder{}, http.StatusNotFound, err
	}

	number, err := ou.OrderPostgresRepository.DeleteDishFromOrder(ctx, orderId, dishId)
	if err != nil {
		return models.DishOrder{}, http.StatusInternalServerError, err
	}

	dishes := models.DishOrder{
		Id:     dishId,
		Number: number,
	}

	return dishes, http.StatusOK, nil
}
