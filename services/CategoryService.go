package services

import (
	"wchan/booking-service/interfaces"
	"wchan/booking-service/models"
	"wchan/calendar-userservice/util"
)

// CategoryService depends on a ICategoryDAO interface
type CategoryService struct {
	DAO interfaces.ICategoryDAO
}

// Create a Category
func (service *CategoryService) Create(c models.Category) error {
	if c.Name == "" {
		return util.ErrEmptyField
	}
	return service.DAO.Create(&c)
}

// Update a category
func (service *CategoryService) Update(c models.Category) error {
	if c.Name == "" {
		return util.ErrEmptyField
	}
	return service.DAO.Update(&c)
}

// Delete a category
func (service *CategoryService) Delete(c models.Category) error {
	if c.Name == "" {
		return util.ErrEmptyField
	}
	// TODO Check that no orgs use this category before deleting
	return service.DAO.Delete(&c)
}

// Find Categories with matching name
func (service *CategoryService) Find(name string) ([]models.Category, error) {
	if name == "" {
		return make([]models.Category, 0), util.ErrEmptyField
	}
	categories, err := service.DAO.Find(name)
	return categories, err
}
