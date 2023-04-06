package service

import (
	"errors"

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
