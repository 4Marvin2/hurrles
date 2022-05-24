package delivery

import (
	"hurrles/internal/models"
	"hurrles/internal/order/usecase"
	"hurrles/internal/pkg/ioutils"
	p "hurrles/internal/pkg/permissions"
	"hurrles/internal/pkg/restaurant_admin"
	userUsecase "hurrles/internal/user/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type OrderHandler struct {
	OrderUseCase usecase.IOrderUsecase
}

func SetOrderRouting(router *mux.Router, os usecase.IOrderUsecase, us userUsecase.IUserUsecase, ss userUsecase.ISessionUsecase) {
	orderHandler := &OrderHandler{
		OrderUseCase: os,
	}

	perm := p.Permission{
		UserUseCase:    us,
		SessionUseCase: ss,
	}

	router.HandleFunc("/api/v1/orders", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(orderHandler.OrdersPost)))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/order", p.SetCSRF(perm.CheckAuth(perm.GetCurrentUser(orderHandler.OrderUserIdGet)))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/order/restaurant", p.SetCSRF(perm.CheckAuth(perm.GetCurrentUser(restaurant_admin.CheckRestaurantAdmin(orderHandler.OrderRestaurantIdGet))))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/order/{id:[0-9]+}", p.SetCSRF(perm.CheckAuth(perm.GetCurrentUser(orderHandler.OrderUserIdDelete)))).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/v1/order/{order_id:[0-9]+}/dish/{dish_id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(orderHandler.DishIdPost))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/order/{order_id:[0-9]+}/dish/{dish_id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(orderHandler.DishIdDelete))).Methods("DELETE", "OPTIONS")
}

func (oh *OrderHandler) OrdersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var order models.Order
	err := ioutils.ReadJSON(r, &order)
	if err != nil {
		log.Errorf("OrderDelivery.OrderPost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	createdOrder, status, err := oh.OrderUseCase.OrdersPost(r.Context(), order)
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.OrderPost: failed create order [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, createdOrder)
}

func (oh *OrderHandler) OrderUserIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	orders, status, err := oh.OrderUseCase.OrderUserIdGet(r.Context())
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.OrderUserIdGet: failed get user's orders [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, orders)
}

func (oh *OrderHandler) OrderUserIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	orderId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("OrderDelivery.OrderUserIdDelete: failed get order id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	deletedOrder, status, err := oh.OrderUseCase.OrderUserIdDelete(r.Context(), uint64(orderId))
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.OrderUserIdDelete: failed delete user's order [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, deletedOrder)
}

func (oh *OrderHandler) DishIdPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	orderId, err := strconv.Atoi(mux.Vars(r)["order_id"])
	if err != nil {
		log.Errorf("OrderDelivery.DishIdPost: failed get order id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	dishId, err := strconv.Atoi(mux.Vars(r)["dish_id"])
	if err != nil {
		log.Errorf("OrderDelivery.DishIdPost: failed get dish id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	dishesNumber, status, err := oh.OrderUseCase.DishIdPost(r.Context(), uint64(dishId), uint64(orderId))
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.DishIdPost: failed add dish to order [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, dishesNumber)
}

func (oh *OrderHandler) DishIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	orderId, err := strconv.Atoi(mux.Vars(r)["order_id"])
	if err != nil {
		log.Errorf("OrderDelivery.DishIdPost: failed get order id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	dishId, err := strconv.Atoi(mux.Vars(r)["dish_id"])
	if err != nil {
		log.Errorf("OrderDelivery.DishIdDelete: failed get dish id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	dishesNumber, status, err := oh.OrderUseCase.DishIdDelete(r.Context(), uint64(dishId), uint64(orderId))
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.DishIdDelete: failed delete dish from order [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, dishesNumber)
}

func (oh *OrderHandler) OrderRestaurantIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	orders, status, err := oh.OrderUseCase.OrderRestaurantIdGet(r.Context())
	if err != nil || status != http.StatusOK {
		log.Errorf("OrderDelivery.OrderRestaurantIdGet: failed get restaurants orders [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, orders)
}
