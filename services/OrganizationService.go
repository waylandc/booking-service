package services

import (
	"wchan/booking-service/interfaces"
	"wchan/booking-service/models"
	"wchan/calendar-userservice/util"
)

// OrganizationService depends on a IOrganizationDAO interface
type OrganizationService struct {
	DAO interfaces.IOrganizationDAO
}

// Create a Organization
func (service *OrganizationService) Create(o models.Organization) error {
	err := validate(o)
	if err != nil {
		return err
	}
	return service.DAO.Create(&o)
}

// Update a Organization
func (service *OrganizationService) Update(o models.Organization) error {
	err := validate(o)
	if err != nil {
		return err
	}

	return service.DAO.Update(&o)
}

// Delete a Organization
func (service *OrganizationService) Delete(o models.Organization) error {
	err := validate(o)
	if err != nil {
		return err
	}
	// TODO Check that no open bookings before deleting
	return service.DAO.Delete(&o)
}

// Find Organizations with matching name
func (service *OrganizationService) Find(name string) ([]models.Organization, error) {
	if name == "" {
		return make([]models.Organization, 0), util.ErrEmptyField
	}
	organizations, err := service.DAO.Find(name)
	return organizations, err
}

func validate(o models.Organization) error {
	if o.Name == "" {
		return util.ErrEmptyField
	}

	if o.Phone == "" {
		return util.ErrEmptyField
	}

	if o.City == "" {
		return util.ErrEmptyField
	}
	return nil
}
