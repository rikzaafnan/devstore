package main

import (
	"fmt"
	"log"
	"net/http"

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
	// create articles
	r.POST("/categories", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "create categories"})
	})

	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(*categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	// get list artcles
	r.GET("/categories", categoryController.BrowseCategory)
	// get detail artcles
	r.GET("/categories/:categoryID", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "create categories"})
	})
	// update artcles
	r.PUT("/categories", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "create categories"})
	})
	// delete detail artcles
	r.DELETE("/categories", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "create categories"})
	})

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)

}
