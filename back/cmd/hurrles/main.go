package main

import (
	"hurrles/config"
	"hurrles/internal/user/delivery"
	"hurrles/internal/user/repository"
	"hurrles/internal/user/usecase"
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
	sessionRepo, err := repository.NewTarantoolConnection(config.Tarantool)
	if err != nil {
		log.Fatal(err)
	}
	userRepo, err := repository.NewPostgresUserRepository(config.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	// usecase
	sessionUseCase := usecase.NewSessionUsecase(sessionRepo, timeoutContext)
	userUseCase := usecase.NewUserUsecase(userRepo, timeoutContext)

	// delivery
	delivery.SetUserRouting(router, userUseCase, sessionUseCase)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.Server.HttpPort,
		WriteTimeout: http.DefaultClient.Timeout,
		ReadTimeout:  http.DefaultClient.Timeout,
	}

	log.Fatal(srv.ListenAndServe())

	log.Infof("STD starting server at %s\n", srv.Addr)
}
