package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

	fmt.Println(cfg.DBDriver)
	fmt.Println(cfg.DBConnection)
	fmt.Println(cfg.ServerPort)

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pinng",
		})
	})

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)

}
