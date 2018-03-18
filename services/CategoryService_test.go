package services

import (
	"testing"
	"wchan/booking-service/models"

	"github.com/stretchr/testify/require"
)

type CategoryDAOStub struct {
}

func (s CategoryDAOStub) Create(c *models.Category) error {
	return nil
}

func (s CategoryDAOStub) Update(c *models.Category) error {
	return nil
}

func (s CategoryDAOStub) Delete(c *models.Category) error {
	return nil
}

func (s CategoryDAOStub) Find(n string) ([]models.Category, error) {
	return make([]models.Category, 0), nil
}

var testCategories = []struct {
	c       models.Category
	success bool
}{
	{models.Category{Name: "Restaurants"}, true},
	{models.Category{Name: ""}, false},
}

func TestCreateCategory(t *testing.T) {
	service := CategoryService{DAO: CategoryDAOStub{}}
	for _, tt := range testCategories {
		err := service.Create(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestUpdateCategory(t *testing.T) {
	service := CategoryService{DAO: CategoryDAOStub{}}
	for _, tt := range testCategories {
		err := service.Update(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestDeleteCategory(t *testing.T) {
	service := CategoryService{DAO: CategoryDAOStub{}}
	for _, tt := range testCategories {
		err := service.Delete(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestFindCategory(t *testing.T) {
	service := CategoryService{DAO: CategoryDAOStub{}}
	categories, err := service.Find("does not exist")
	require.Nil(t, err)
	require.Empty(t, categories)

	categories, err = service.Find("")
	require.NotNil(t, err)
}
