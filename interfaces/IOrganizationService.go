package interfaces

import "wchan/booking-service/models"

// IOrganizationService interface
type IOrganizationService interface {
	Create(c models.Organization) error
	Update(c models.Organization) error
	Delete(c models.Organization) error
	Find(name string) ([]models.Organization, error)
}
