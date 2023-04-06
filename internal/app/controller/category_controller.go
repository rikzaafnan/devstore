package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rikzaafnan/devstore/internal/app/service"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {

	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
