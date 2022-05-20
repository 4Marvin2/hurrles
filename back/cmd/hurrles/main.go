package main

import (
	"hurrles/config"
	admin_delivery "hurrles/internal/admin/delivery"
	admin_repository "hurrles/internal/admin/repository"
	admin_usecase "hurrles/internal/admin/usecase"

	restaurant_delivery "hurrles/internal/restaurant/delivery"
	restaurant_repository "hurrles/internal/restaurant/repository"
	restaurant_usecase "hurrles/internal/restaurant/usecase"

	user_delivery "hurrles/internal/user/delivery"
	user_repository "hurrles/internal/user/repository"
	user_usecase "hurrles/internal/user/usecase"

	order_delivery "hurrles/internal/order/delivery"
	order_repository "hurrles/internal/order/repository"
	order_usecase "hurrles/internal/order/usecase"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.SetConfig()

	// router
	router := mux.NewRouter()

	timeoutContext := config.Timeouts.ContextTimeout

	// repository
	sessionRepo, err := user_repository.NewTarantoolConnection(config.Tarantool)
	if err != nil {
		log.Fatal(err)
	}
	userRepo, err := user_repository.NewPostgresUserRepository(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	adminRepo, err := admin_repository.NewPostgresAdminRepository(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	restaurantRepo, err := restaurant_repository.NewPostgresRestaurantRepository(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	orderRepo, err := order_repository.NewPostgresOrderRepository(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	// usecase
	sessionUseCase := user_usecase.NewSessionUsecase(sessionRepo, timeoutContext)
	userUseCase := user_usecase.NewUserUsecase(userRepo, timeoutContext)
	adminUseCase := admin_usecase.NewAdminUsecase(adminRepo, timeoutContext)
	restaurantUseCase := restaurant_usecase.NewRestaurantUsecase(restaurantRepo, timeoutContext)
	orderUseCase := order_usecase.NewOrderUsecase(orderRepo, timeoutContext)

	// delivery
	user_delivery.SetUserRouting(router, userUseCase, sessionUseCase)
	admin_delivery.SetAdminRouting(router, adminUseCase, userUseCase, sessionUseCase)
	restaurant_delivery.SetRestaurantRouting(router, restaurantUseCase, userUseCase, sessionUseCase)
	order_delivery.SetOrderRouting(router, orderUseCase, userUseCase, sessionUseCase)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.Server.HttpPort,
		WriteTimeout: http.DefaultClient.Timeout,
		ReadTimeout:  http.DefaultClient.Timeout,
	}

	log.Fatal(srv.ListenAndServe())

	log.Infof("STD starting server at %s\n", srv.Addr)
}
