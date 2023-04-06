package service

import "github.com/rikzaafnan/devstore/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
}
