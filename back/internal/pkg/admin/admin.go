package admin

import (
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/pkg/ioutils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CheckAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		curUser, ok := r.Context().Value(config.ContextUser).(models.User)
		if !ok {
			log.Errorf("Admin.CheckAdmin: no ctx user")
			ioutils.SendError(w, http.StatusForbidden, "")
			return
		}

		if curUser.IsAdmin {
			next.ServeHTTP(w, r)
			return
		}

		log.Errorf("Admin.CheckAdmin: user isn't admin")
		ioutils.SendError(w, http.StatusForbidden, "")
		return
	})
}
