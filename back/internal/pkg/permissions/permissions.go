package permissions

import (
	"context"
	"hurrles/config"
	"hurrles/internal/pkg/ioutils"
	"hurrles/internal/user/models"
	"hurrles/internal/user/usecase"

	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

type Permission struct {
	UserUseCase    usecase.IUserUsecase
	SessionUseCase usecase.ISessionUsecase
}

func (perm *Permission) CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userSession models.Session
		session, err := r.Cookie("sessionId")
		if err != nil {
			log.Errorf("Permissions.CheckAuth: no cookie: %w", err)
			ioutils.SendError(w, http.StatusForbidden, "")
			return
		} else {
			userSession, err = perm.SessionUseCase.GetSessionByCookie(r.Context(), session.Value)
			if err != nil {
				log.Errorf("Permissions.CheckAuth: failed GetSessionByCookie with error: %w", err)
				ioutils.SendError(w, http.StatusForbidden, "")
				return
			}
		}
		r = r.WithContext(context.WithValue(r.Context(), config.ContextUserID, userSession))
		next.ServeHTTP(w, r)
	})
}

func (perm *Permission) GetCurrentUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxSession := r.Context().Value(config.ContextUserID)
		if ctxSession == nil {
			log.Errorf("Permissions.GetCurrentUser: empty session in ctx")
			ioutils.SendError(w, http.StatusForbidden, "")
			return
		}
		currentSession, ok := ctxSession.(models.Session)
		if !ok {
			log.Errorf("Permissions.GetCurrentUser: failed cast ctxSession to Session")
			ioutils.SendError(w, http.StatusForbidden, "")
			return
		}

		currentUser, status, err := perm.UserUseCase.GetUserById(r.Context(), currentSession.UserID)
		if err != nil || status != http.StatusOK {
			log.Errorf("Permissions.GetCurrentUser: failed GetUserById with [error: %w] [status: %d]", err, status)
			ioutils.SendError(w, status, "")
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), config.ContextUser, currentUser))
		next.ServeHTTP(w, r)
	})
}

func generateCsrfLogic(w http.ResponseWriter) {
	csrf, err := uuid.NewV4()
	if err != nil {
		ioutils.SendError(w, http.StatusForbidden, "")
		return
	}
	timeDelta := time.Now().Add(time.Hour * 24 * 30)
	csrfCookie := &http.Cookie{Name: "csrf", Value: csrf.String(), Path: "/", HttpOnly: true, Expires: timeDelta}

	http.SetCookie(w, csrfCookie)
	w.Header().Set("csrf", csrf.String())
}

func SetCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})
}

func CheckCSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			csrf := r.Header.Get("x-csrf-Token")
			csrfCookie, err := r.Cookie("csrf")

			if err != nil || csrf == "" || csrfCookie.Value == "" || csrfCookie.Value != csrf {
				ioutils.SendError(w, http.StatusForbidden, "")
				return
			}
			generateCsrfLogic(w)
			next.ServeHTTP(w, r)
		})
}
