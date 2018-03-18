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

func (dao CategoryDAOStub) Delete(c *models.Category) error {
	return nil
}

func (dao CategoryDAOStub) Update(c *models.Category) error {
	return nil
}

func (dao CategoryDAOStub) Find(name string) ([]models.Category, error) {
	return make([]models.Category, 0), nil
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

func TestCreateBadRequest(t *testing.T) {
	values := map[string]string{"wrongfield": "should fail"}
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
