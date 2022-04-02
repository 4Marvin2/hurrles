package delivery

import (
	"hurrles/internal/pkg/ioutils"
	p "hurrles/internal/pkg/permissions"
	"hurrles/internal/user/models"
	"hurrles/internal/user/usecase"
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUCase usecase.IUserUsecase
}

func SetUserRouting(router *mux.Router, us usecase.IUserUsecase, ss usecase.ISessionUsecase) {
	userHandler := &UserHandler{
		UserUCase: us,
	}

	perm := p.Permission{
		UserUseCase:    us,
		SessionUseCase: ss,
	}

	router.HandleFunc("/api/v1/user/signup", p.SetCSRF(userHandler.UserSignupPost)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/user/login", p.SetCSRF(userHandler.UserLoginPost)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/user/logout", perm.CheckAuth(userHandler.UserLogoutGet)).Methods("GET", "OPTIONS")
}

func (uh *UserHandler) UserLoginPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var credentials models.UserCredentials
	err := ioutils.ReadJSON(r, &credentials)
	if err != nil {
		log.Errorf("UserDelivery.UserLoginPost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
	}

	user, status, err := uh.UserUCase.LoginUser(credentials)
	if err != nil || status != http.StatusOK {
		log.Errorf("UserDelivery.UserLoginPost: failed user verification with [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
	}

	// cookie

	ioutils.Send(w, status, user)
}

func (uh *UserHandler) UserSignupPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var user models.User
	err := ioutils.ReadJSON(r, &user)
	if err != nil {
		log.Errorf("UserDelivery.UserSignupPost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
	}

	createdUser, status, err := uh.UserUCase.SignupUser(user)
	if err != nil || status != http.StatusOK {
		log.Errorf("UserDelivery.UserSignupPost: failed user registration with [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
	}

	// cookie

	ioutils.Send(w, status, createdUser)
}

func (uh *UserHandler) UserLogoutGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
