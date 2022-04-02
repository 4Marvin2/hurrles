package delivery

import (
	"hurrles/internal/user"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUCase user.IUserUsecase
}

func SetUserRouting(router *mux.Router, us user.IUserUsecase) {
	userHandler := &UserHandler{
		UserUCase: us,
	}

	router.HandleFunc("/api/v1/user", userHandler.UserLoginPost).Methods("POST", "OPTIONS")
}

func (uh *UserHandler) UserLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) UserLogoutGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) UserSignupPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
