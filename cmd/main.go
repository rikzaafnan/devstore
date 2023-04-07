package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rikzaafnan/devstore/internal/app/controller"
	"github.com/rikzaafnan/devstore/internal/app/repository"
	"github.com/rikzaafnan/devstore/internal/app/service"
	"github.com/rikzaafnan/devstore/internal/pkg/config"
	"github.com/rikzaafnan/devstore/internal/pkg/db"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	// read configuration
	configload, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot log app config")
	}

	cfg = configload

	// connect database
	db, err := db.Connection(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db not established")
	}

	DBConn = db
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pinng",
		})
	})

	// endpoint

	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(*categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	// get list artcles
	r.GET("/categories", categoryController.BrowseCategory)
	// create articles
	r.POST("/categories", categoryController.CreateCategory)
	// get detail artcles
	r.GET("/categories/:categoryID", categoryController.DetailCategory)
	// update artcles
	r.PUT("/categories/:categoryID", categoryController.UpdateCategory)
	// delete detail artcles
	r.DELETE("/categories/:categoryID", categoryController.DeleteCategory)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)

}
