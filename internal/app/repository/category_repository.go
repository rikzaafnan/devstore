package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rikzaafnan/devstore/internal/app/model"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) Create() {

}

func (cr *CategoryRepository) Browse() ([]model.Category, error) {

	var (
		categories   []model.Category
		sqlStatement = `SELECT id, name, description FROM categories`
	)

	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Println(fmt.Errorf("error CategoryRepository - Browse : %s", err))
		return categories, err
	}

	for rows.Next() {

		var category model.Category
		rows.StructScan(&category)
		categories = append(categories, category)

	}

	return categories, nil

}

// func (cr *CategoryRepository) Browse() {

// }

// func (cr *CategoryRepository) Browse() {

// }

// func (cr *CategoryRepository) Browse() {

// }
