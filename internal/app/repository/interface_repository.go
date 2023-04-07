package repository

import "github.com/rikzaafnan/devstore/internal/app/model"

type ICategoryRepository interface {
	Browse() ([]model.Category, error)
	Create(categoory model.Category) error
	FindOneByID(categoryID int) (model.Category, error)
	Update(categoryID int, req model.Category) (lastInsertID int, err error)
	DeleteByID(categoryID int) error
}
