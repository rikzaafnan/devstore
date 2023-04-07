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

func (cr *CategoryRepository) Create(category model.Category) error {

	var (
		sqlStatement = `INSERT INTO categories (name, description) VALUES ($1, $2)`
	)

	_, err := cr.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		log.Println(fmt.Errorf("error CategoryRepository - Create : %s", err))
		return err
	}

	return nil

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

func (cr *CategoryRepository) FindOneByID(categoryID int) (model.Category, error) {

	var (
		category     model.Category
		sqlStatement = `SELECT id, name, description FROM categories where id = $1`
	)

	row := cr.DB.QueryRowx(sqlStatement, categoryID)

	err := row.StructScan(&category)
	if err != nil {
		log.Println(fmt.Errorf("error CategoryRepository - FindOneByID : %s", err))
		return category, err
	}

	return category, nil

}

func (cr *CategoryRepository) Update(categoryID int, req model.Category) (lastInsertID int, err error) {

	var (
		sqlStatement = `UPDATE categories
		SET name = $1, description = $2 
		WHERE id = $3 
		RETURNING id`
	)

	var id int
	err = cr.DB.QueryRow(sqlStatement, req.Name, req.Description, categoryID).Scan(&id)
	if err != nil {
		log.Println(fmt.Errorf("error CategoryRepository - Update : %s", err))
		return 0, err
	}

	// lastInsertIDDB, err := rows.LastInsertId()
	// if err != nil {
	// 	log.Println(fmt.Errorf("error CategoryRepository - Update : %s", err))
	// 	return 0, err
	// }

	if id == 0 {
		log.Println(fmt.Errorf("error CategoryRepository - Update : %s", "error lastinsertID"))
		return 0, err
	}

	lastInsertID = id

	return lastInsertID, nil

}

func (cr *CategoryRepository) DeleteByID(categoryID int) error {

	var (
		sqlStatement = `DELETE FROM categories
		WHERE id = $1`
	)

	_, err := cr.DB.Exec(sqlStatement, categoryID)
	if err != nil {
		log.Println(fmt.Errorf("error CategoryRepository - Update : %s", err))
		return err
	}

	return nil

}
