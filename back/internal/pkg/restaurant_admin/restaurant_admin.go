package restaurant_admin

import (
	"hurrles/config"
	"hurrles/internal/models"
	"hurrles/internal/pkg/ioutils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CheckRestaurantAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		curUser, ok := r.Context().Value(config.ContextUser).(models.User)
		if !ok {
			log.Errorf("RestaurantAdmin.CheckRestaurantAdmin: no ctx user")
			ioutils.SendError(w, http.StatusForbidden, "")
			return
		}

		if curUser.IsAdmin {
			next.ServeHTTP(w, r)
			return
		}

		log.Errorf("RestaurantAdmin.CheckRestaurantAdmin: user isn't restaurant admin")
		ioutils.SendError(w, http.StatusForbidden, "")
		return
	})
}
