package delivery

import (
	"hurrles/internal/models"
	"hurrles/internal/pkg/ioutils"
	p "hurrles/internal/pkg/permissions"
	"hurrles/internal/restaurant/usecase"
	userUsecase "hurrles/internal/user/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type RestaurantHandler struct {
	RestaurantUseCase usecase.IRestaurantUsecase
}

func SetRestaurantRouting(router *mux.Router, rs usecase.IRestaurantUsecase, us userUsecase.IUserUsecase, ss userUsecase.ISessionUsecase) {
	restaurantHandler := &RestaurantHandler{
		RestaurantUseCase: rs,
	}

	perm := p.Permission{
		UserUseCase:    us,
		SessionUseCase: ss,
	}

	router.HandleFunc("/api/v1/restaurants", p.SetCSRF(perm.CheckAuth(restaurantHandler.RestaurantsGet))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/{id:[0-9]+}/menu", p.SetCSRF(perm.CheckAuth(restaurantHandler.RestaurantIdMenuGet))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/{id:[0-9]+}/places", p.CheckCSRF(perm.CheckAuth(restaurantHandler.RestaurantIdPlacesGet))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/{id:[0-9]+}/places", p.CheckCSRF(perm.CheckAuth(restaurantHandler.RestaurantIdAllPlacesGet))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/favorite/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(restaurantHandler.RestaurantFavoritePost)))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/favorite", p.SetCSRF(perm.CheckAuth(perm.GetCurrentUser(restaurantHandler.RestaurantFavoriteGet)))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/favorite/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(restaurantHandler.RestaurantFavoriteDelete)))).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/v1/restaurant/search", p.CheckCSRF(perm.CheckAuth(restaurantHandler.RestaurantSearchPost))).Methods("POST", "OPTIONS")
}

func (rh *RestaurantHandler) RestaurantIdMenuGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdMenuGet: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	dishes, status, err := rh.RestaurantUseCase.RestaurantIdMenuGet(r.Context(), uint64(restaurantId))
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantIdMenuGet: failed get menu [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, dishes)
}

func (rh *RestaurantHandler) RestaurantIdPlacesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	var placeParameters models.PlaceParameters
	err = ioutils.ReadJSON(r, &placeParameters)
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}
	placeParameters.RestaurantId = uint64(restaurantId)

	places, status, err := rh.RestaurantUseCase.RestaurantIdPlacesGet(r.Context(), placeParameters)
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get places [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, places)
}

func (rh *RestaurantHandler) RestaurantIdAllPlacesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	places, status, err := rh.RestaurantUseCase.RestaurantIdAllPlacesGet(r.Context(), uint64(restaurantId))
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get places [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, places)
}

func (rh *RestaurantHandler) RestaurantsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurants, status, err := rh.RestaurantUseCase.RestaurantsGet(r.Context())
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantsGet: failed get restaurants [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, restaurants)
}

func (rh *RestaurantHandler) RestaurantFavoritePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	status, err := rh.RestaurantUseCase.RestaurantFavoritePost(r.Context(), uint64(restaurantId))
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantsFavoritePost: failed add favorite restaurant [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.SendWithoutBody(w, status)
}

func (rh *RestaurantHandler) RestaurantFavoriteGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurants, status, err := rh.RestaurantUseCase.RestaurantFavoriteGet(r.Context())
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantsGet: failed get favorite restaurants [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, restaurants)
}

func (rh *RestaurantHandler) RestaurantFavoriteDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantIdPlacesGet: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	status, err := rh.RestaurantUseCase.RestaurantFavoriteDelete(r.Context(), uint64(restaurantId))
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantsFavoritePost: failed delete favorite restaurant [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.SendWithoutBody(w, status)
}

func (rh *RestaurantHandler) RestaurantSearchPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var search models.Search
	err := ioutils.ReadJSON(r, &search)
	if err != nil {
		log.Errorf("RestaurantDelivery.RestaurantSearchPost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	restaurants, status, err := rh.RestaurantUseCase.RestaurantSearchPost(r.Context(), search.SearchPattern)
	if err != nil || status != http.StatusOK {
		log.Errorf("RestaurantDelivery.RestaurantSearchPost: failed search restaurants by string=%s with [error: %w] [status: %d]", search.SearchPattern, err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, restaurants)
}
