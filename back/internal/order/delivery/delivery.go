package delivery

import (
	"hurrles/internal/models"
	"hurrles/internal/order/usecase"
	"hurrles/internal/pkg/ioutils"
	p "hurrles/internal/pkg/permissions"
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
	router.HandleFunc("/api/v1/order/{id:[0-9]+}", p.SetCSRF(perm.CheckAuth(perm.GetCurrentUser(orderHandler.OrderUserIdDelete)))).Methods("DELETE", "OPTIONS")
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
