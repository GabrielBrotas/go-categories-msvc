package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
	c_use_cases "github.com/GabrielBrotas/go-categories-msvc/internal/categories/use-cases"
)

func getCategories(c *gin.Context, repository *c_repository.CategoryRepository) {
	use_case := c_use_cases.NewGetCategoriesUseCase(repository)
	categories, err := use_case.Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "result": categories})

}
