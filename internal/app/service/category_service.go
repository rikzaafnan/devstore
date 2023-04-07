package service

import (
	"errors"

	"github.com/rikzaafnan/devstore/internal/app/model"
	"github.com/rikzaafnan/devstore/internal/app/repository"
	"github.com/rikzaafnan/devstore/internal/app/schema"
)

type CategoryService struct {
	repository repository.ICategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {

	return &CategoryService{repository: &repo}

}

func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {

	var resp []schema.GetCategoryResp

	categories, err := cs.repository.Browse()
	if err != nil {
		return resp, errors.New("cannot get categories")
	}

	for _, v := range categories {
		var respData schema.GetCategoryResp
		respData.ID = v.ID
		respData.Name = v.Name
		respData.Description = v.Description

		resp = append(resp, respData)
	}

	return resp, nil

}

func (cs *CategoryService) Create(req schema.CreateCategoryReq) error {

	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repository.Create(insertData)
	if err != nil {
		return err
	}

	return nil

}

func (cs *CategoryService) FindOneByID(categoryID int) (schema.GetCategoryResp, error) {

	var category schema.GetCategoryResp

	categoryModel, err := cs.repository.FindOneByID(categoryID)
	if err != nil {
		return category, err
	}

	category.ID = categoryModel.ID
	category.Name = categoryModel.Name
	category.Description = categoryModel.Description

	return category, nil

}

func (cs *CategoryService) Update(categoryID int, req schema.CreateCategoryReq) (schema.GetCategoryResp, error) {

	var categorySchema schema.GetCategoryResp

	// find by id
	category, err := cs.repository.FindOneByID(categoryID)
	if err != nil {
		return categorySchema, err
	}

	category.Description = req.Description
	category.Name = req.Name

	lastInsertIDCategory, err := cs.repository.Update(categoryID, category)
	if err != nil {
		return categorySchema, err
	}

	// find by id
	category, err = cs.repository.FindOneByID(lastInsertIDCategory)
	if err != nil {
		return categorySchema, err
	}

	categorySchema.ID = category.ID
	categorySchema.Name = category.Name
	categorySchema.Description = category.Description

	return categorySchema, nil

}

func (cs *CategoryService) DeleteByID(categoryID int) error {

	// find by id
	category, err := cs.repository.FindOneByID(categoryID)
	if err != nil {
		return err
	}

	err = cs.repository.DeleteByID(category.ID)
	if err != nil {
		return err
	}

	return nil

}
