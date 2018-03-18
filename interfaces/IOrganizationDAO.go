package interfaces

import "wchan/booking-service/models"

// IOrganizationDAO is our dao for persisting Organizations
type IOrganizationDAO interface {
	Create(c *models.Organization) error
	Delete(c *models.Organization) error
	Update(c *models.Organization) error
	Find(name string) ([]models.Organization, error)
}
