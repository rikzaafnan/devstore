package repository

import "github.com/rikzaafnan/devstore/internal/app/model"

type ICategoryRepository interface {
	Browse() ([]model.Category, error)
}
