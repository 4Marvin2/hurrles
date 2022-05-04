package delivery

import (
	"crypto/md5"
	"fmt"
	"hurrles/internal/models"
	"hurrles/internal/pkg/hasher"
	"hurrles/internal/pkg/ioutils"
	"net/http"
	"time"

	p "hurrles/internal/pkg/permissions"
	"hurrles/internal/user/usecase"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUseCase    usecase.IUserUsecase
	SessionUseCase usecase.ISessionUsecase
}

func SetUserRouting(router *mux.Router, us usecase.IUserUsecase, ss usecase.ISessionUsecase) {
	userHandler := &UserHandler{
		UserUseCase:    us,
		SessionUseCase: ss,
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
		return
	}

	user, status, err := uh.UserUseCase.LoginUser(r.Context(), credentials)
	if err != nil || status != http.StatusOK {
		log.Errorf("UserDelivery.UserLoginPost: failed user verification with [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	cookie, err := uh.createSessionCookie(user.Email)
	if err != nil {
		log.Errorf("UserDelivery.UserLoginPost: failed create cookie for user with error: %w", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	sess := models.Session{
		Cookie: cookie.Value,
		UserID: user.Id,
	}

	err = uh.SessionUseCase.AddSession(r.Context(), sess)
	if err != nil {
		log.Errorf("UserDelivery.UserLoginPost: failed add session in tnt for user with error: %w", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	http.SetCookie(w, &cookie)

	ioutils.Send(w, status, user)
}

func (uh *UserHandler) UserSignupPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var user models.User
	err := ioutils.ReadJSON(r, &user)
	if err != nil {
		log.Errorf("UserDelivery.UserSignupPost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	createdUser, status, err := uh.UserUseCase.SignupUser(r.Context(), user)
	if err != nil || status != http.StatusOK {
		log.Errorf("UserDelivery.UserSignupPost: failed user registration with [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	cookie, err := uh.createSessionCookie(user.Email)
	if err != nil {
		log.Errorf("UserDelivery.UserSignupPost: failed create cookie for user with error: %w", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	sess := models.Session{
		Cookie: cookie.Value,
		UserID: createdUser.Id,
	}

	err = uh.SessionUseCase.AddSession(r.Context(), sess)
	if err != nil {
		log.Errorf("UserDelivery.UserSignupPost: failed add session in tnt for user with error: %w", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	http.SetCookie(w, &cookie)

	ioutils.Send(w, status, createdUser)
}

func (uh *UserHandler) UserLogoutGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := uh.SessionUseCase.DeleteSession(r.Context())
	if err != nil {
		log.Errorf("UserDelivery.UserLogoutGet: failed delete session in tnt for user with error: %w", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	authCookie := &http.Cookie{
		Name:     "sessionId",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, authCookie)

	csrfCookie := &http.Cookie{
		Name:     "csrf",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, csrfCookie)

	ioutils.SendWithoutBody(w, http.StatusOK)
}

func (uh *UserHandler) createSessionCookie(email string) (http.Cookie, error) {
	expiration := time.Now().Add(10 * time.Hour)

	hashedEmail, err := hasher.HashAndSalt(email)
	if err != nil {
		return http.Cookie{}, err
	}
	data := hashedEmail + time.Now().String()
	md5CookieValue := fmt.Sprintf("%x", md5.Sum([]byte(data)))

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    md5CookieValue,
		Expires:  expiration,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/api/v1",
	}

	return cookie, nil
}
