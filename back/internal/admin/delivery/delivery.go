package delivery

import (
	"hurrles/internal/admin/usecase"
	"hurrles/internal/models"
	"hurrles/internal/pkg/admin"
	"hurrles/internal/pkg/ioutils"
	p "hurrles/internal/pkg/permissions"
	userUsecase "hurrles/internal/user/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type AdminHandler struct {
	AdminUseCase usecase.IAdminUsecase
}

func SetAdminRouting(router *mux.Router, as usecase.IAdminUsecase, us userUsecase.IUserUsecase, ss userUsecase.ISessionUsecase) {
	adminHandler := &AdminHandler{
		AdminUseCase: as,
	}

	perm := p.Permission{
		UserUseCase:    us,
		SessionUseCase: ss,
	}

	router.HandleFunc("/api/v1/admin/restaurant", p.CheckCSRF((perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.RestaurantCreatePost)))))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/admin/restaurant/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.RestaurantUpdatePut))))).Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/v1/admin/dish", p.CheckCSRF((perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.DishCreatePost)))))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/admin/dish/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.DishUpdatePut))))).Methods("PUT", "OPTIONS")

	router.HandleFunc("/api/v1/admin/place", p.CheckCSRF((perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.PlaceCreatePost)))))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/admin/place/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.PlaceUpdatePut))))).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/admin/places", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.AllPlacesUpdatePut))))).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/admin/place/{id:[0-9]+}", p.CheckCSRF(perm.CheckAuth(perm.GetCurrentUser(admin.CheckAdmin(adminHandler.PlaceDropDelete))))).Methods("DELETE", "OPTIONS")
}

//
// RESTAURANTS
//

func (ah *AdminHandler) RestaurantCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var restaurant models.Restaurant
	err := ioutils.ReadJSON(r, &restaurant)
	if err != nil {
		log.Errorf("AdminDelivery.RestaurantCreatePost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	createdRestaurant, status, err := ah.AdminUseCase.RestaurantCreatePost(r.Context(), restaurant)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.RestaurantCreatePost: failed create restaurant [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, createdRestaurant)
}

func (rh *AdminHandler) RestaurantUpdatePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	restaurantId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("AdminDelivery.RestaurantUpdatePut: failed get restaurant id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	var restaurant models.Restaurant
	err = ioutils.ReadJSON(r, &restaurant)
	if err != nil {
		log.Errorf("AdminDelivery.RestaurantUpdatePut: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}
	restaurant.Id = uint64(restaurantId)

	updatedRestaurant, status, err := rh.AdminUseCase.RestaurantUpdatePut(r.Context(), restaurant)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.RestaurantUpdatePut: failed update restaurant [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, updatedRestaurant)
}

//
// DISHES
//

func (ah *AdminHandler) DishCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var dish models.Dish
	err := ioutils.ReadJSON(r, &dish)
	if err != nil {
		log.Errorf("AdminDelivery.DishCreatePost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	createdDish, status, err := ah.AdminUseCase.DishCreatePost(r.Context(), dish)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.DishCreatePost: failed create dish [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, createdDish)
}

func (rh *AdminHandler) DishUpdatePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dishId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("AdminDelivery.DishUpdatePut: failed get dish id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	var dish models.Dish
	err = ioutils.ReadJSON(r, &dish)
	if err != nil {
		log.Errorf("AdminDelivery.DishUpdatePut: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}
	dish.Id = uint64(dishId)

	updatedDish, status, err := rh.AdminUseCase.DishUpdatePut(r.Context(), dish)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.DishUpdatePut: failed update dish [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, updatedDish)
}

//
// PLACES
//

func (ah *AdminHandler) PlaceCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var place models.Place
	err := ioutils.ReadJSON(r, &place)
	if err != nil {
		log.Errorf("AdminDelivery.PlaceCreatePost: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	createdPlace, status, err := ah.AdminUseCase.PlaceCreatePost(r.Context(), place)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.PlaceCreatePost: failed create place [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, createdPlace)
}

func (rh *AdminHandler) PlaceUpdatePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	placeId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed get place id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	var place models.Place
	err = ioutils.ReadJSON(r, &place)
	if err != nil {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}
	place.Id = uint64(placeId)

	updatedDish, status, err := rh.AdminUseCase.PlaceUpdatePut(r.Context(), place)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed update place [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, updatedDish)
}

func (rh *AdminHandler) AllPlacesUpdatePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var places models.PlaceList
	err := ioutils.ReadJSON(r, &places)
	if err != nil {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed read json with error: %w", err)
		ioutils.SendError(w, http.StatusBadRequest, "")
		return
	}

	updatedPlaces, status, err := rh.AdminUseCase.AllPlacesUpdatePut(r.Context(), places)
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed update place [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, updatedPlaces)
}

func (rh *AdminHandler) PlaceDropDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	placeId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed get place id from url [error: %w]", err)
		ioutils.SendError(w, http.StatusInternalServerError, "")
		return
	}

	updatedDish, status, err := rh.AdminUseCase.PlaceDropDelete(r.Context(), uint64(placeId))
	if err != nil || status != http.StatusOK {
		log.Errorf("AdminDelivery.PlaceUpdatePut: failed update place [error: %w] [status: %d]", err, status)
		ioutils.SendError(w, status, "")
		return
	}

	ioutils.Send(w, status, updatedDish)
}
