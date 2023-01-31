package controllers

import (
	"net/http"

	"github.com/GabrielBrotas/go-categories-msvc/internal/repositories"
	use_cases "github.com/GabrielBrotas/go-categories-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
