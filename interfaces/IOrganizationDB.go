package interfaces

import "wchan/booking-service/models"

// IOrganizationDB interface
type IOrganizationDB interface {
	Create(c models.Organization) error
	Update(c models.Organization) error
	Delete(c models.Organization) error
	Find(name string) ([]models.Organization, error)
}
