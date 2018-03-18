package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wchan/booking-service/models"
	"wchan/booking-service/services"

	"github.com/stretchr/testify/assert"
)

type CategoryDAOStub struct {
}

// Create is stubbed out here so we just return nil
func (dao CategoryDAOStub) Create(c *models.Category) error {
	return nil
}

func TestCreate(t *testing.T) {
	values := map[string]string{"name": "Restaurants"}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", "/category", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	service := &services.CategoryService{DAO: &CategoryDAOStub{}}
	handler := CategoryHandler{service}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestCreateWithEmptyField(t *testing.T) {
	values := map[string]string{"name": ""}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", "/category", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	service := &services.CategoryService{DAO: &CategoryDAOStub{}}
	handler := CategoryHandler{service}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)

}
