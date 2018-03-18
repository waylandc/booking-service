package interfaces

import "wchan/booking-service/models"

// ICategoryService interface
type ICategoryService interface {
	Create(c models.Category) error
	Delete(c *models.Category) error
	Update(c *models.Category) error
	Find(name string) ([]models.Category, error)
}
