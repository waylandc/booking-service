package handlers

import (
	"encoding/json"
	"net/http"
	"wchan/booking-service/interfaces"
	"wchan/booking-service/models"
	"wchan/calendar-userservice/util"
)

// CategoryHandler depends on an ICategoryService
type CategoryHandler struct {
	Service interfaces.ICategoryService
}

func (handler *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = handler.Service.Create(newCategory)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(util.HTTPCodeFrom(err))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}
