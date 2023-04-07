package service

import "github.com/rikzaafnan/devstore/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req schema.CreateCategoryReq) error
	FindOneByID(categoryID int) (schema.GetCategoryResp, error)
	Update(categoryID int, req schema.CreateCategoryReq) (schema.GetCategoryResp, error)
	DeleteByID(categoryID int) error
}
