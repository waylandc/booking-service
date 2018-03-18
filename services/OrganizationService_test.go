package services

import (
	"testing"
	"wchan/booking-service/models"

	"github.com/stretchr/testify/require"
)

type OrganizationDAOStub struct {
}

func (s OrganizationDAOStub) Create(c *models.Organization) error {
	return nil
}

func (s OrganizationDAOStub) Update(c *models.Organization) error {
	return nil
}

func (s OrganizationDAOStub) Delete(c *models.Organization) error {
	return nil
}

func (s OrganizationDAOStub) Find(n string) ([]models.Organization, error) {
	return make([]models.Organization, 0), nil
}

var testOrganizations = []struct {
	c       models.Organization
	success bool
}{
	{models.Organization{Name: "The Diner", Website: "www.diner.com", City: "Hong Kong",
		Phone: "9541 8599", Categories: []models.Category{{Name: "Restaurants"}}}, true},
	{models.Organization{Name: "Nan Tei", Website: "www.openrice.com", City: "Hong Kong",
		Phone: "9999 8888", Categories: []models.Category{{Name: "Restaurants"}}}, true},
	{models.Organization{Name: "", Website: "www.openrice.com", City: "Hong Kong",
		Phone: "9999 8888", Categories: []models.Category{{Name: "Restaurants"}}}, false},
	{models.Organization{Name: "Nan Tei", Website: "", City: "",
		Phone: "9999 8888", Categories: []models.Category{{Name: "Restaurants"}}}, false},
	{models.Organization{Name: "Wah sing", Website: "www.yelp.ca", City: "Toronto",
		Phone: "", Categories: []models.Category{{Name: "Restaurants"}}}, false},
}

func TestCreateOrganization(t *testing.T) {
	service := OrganizationService{DAO: OrganizationDAOStub{}}
	for _, tt := range testOrganizations {
		err := service.Create(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestUpdateOrganization(t *testing.T) {
	service := OrganizationService{DAO: OrganizationDAOStub{}}
	for _, tt := range testOrganizations {
		err := service.Update(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestDeleteOrganization(t *testing.T) {
	service := OrganizationService{DAO: OrganizationDAOStub{}}
	for _, tt := range testOrganizations {
		err := service.Delete(tt.c)
		if tt.success {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestFindOrganization(t *testing.T) {
	service := OrganizationService{DAO: OrganizationDAOStub{}}
	organizations, err := service.Find("does not exist")
	require.Nil(t, err)
	require.Empty(t, organizations)

	organizations, err = service.Find("")
	require.NotNil(t, err)
}
