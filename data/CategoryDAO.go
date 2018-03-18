package data

import (
	"wchan/booking-service/interfaces"
	"wchan/booking-service/models"
)

// CategoryDAO has a dependency on ICategoryDB interface
type CategoryDAO struct {
	Repo interfaces.ICategoryDB
}

// Create a category record
func (dao *CategoryDAO) Create(cat *models.Category) error {
	err := dao.Repo.Create(*cat)

	return err
}

// Update a category record
func (dao *CategoryDAO) Update(cat *models.Category) error {
	err := dao.Repo.Update(*cat)

	return err
}

// Delete a category record
func (dao *CategoryDAO) Delete(cat *models.Category) error {
	err := dao.Repo.Delete(*cat)

	return err
}
