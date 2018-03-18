package interfaces

import "wchan/booking-service/models"

// ICategoryDB interface
type ICategoryDB interface {
	Create(c models.Category) error
	Update(c models.Category) error
	Delete(c models.Category) error
	Find(name string) ([]models.Category, error)
}
