package interfaces

import "wchan/booking-service/models"

// ICategoryDAO is our dao for persisting Categories
type ICategoryDAO interface {
	Create(c *models.Category) error
	Delete(c *models.Category) error
	Update(c *models.Category) error
	Find(name string) ([]models.Category, error)
}
