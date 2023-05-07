package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
	c_use_cases "github.com/GabrielBrotas/go-categories-msvc/internal/categories/use-cases"
	utils "github.com/GabrielBrotas/go-categories-msvc/pkg/utils"
)

func deleteCategory(c *gin.Context, repository *c_repository.CategoryRepository) {
	id, err := utils.StringToUint(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	use_case := c_use_cases.NewDeleteCategoryUseCase(repository)
	err = use_case.Execute(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
