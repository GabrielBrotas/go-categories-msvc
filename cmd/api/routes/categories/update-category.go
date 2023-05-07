package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
	c_use_cases "github.com/GabrielBrotas/go-categories-msvc/internal/categories/use-cases"
	utils "github.com/GabrielBrotas/go-categories-msvc/pkg/utils"
)

type updateCategoryBody struct {
	Name string `json:"name"`
}

func updateCategory(c *gin.Context, repository *c_repository.CategoryRepository) {
	id, err := utils.StringToUint(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var requestBody updateCategoryBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	use_case := c_use_cases.NewUpdateCategoryUseCase(repository)
	err = use_case.Execute(c_use_cases.UpdateCategoryInput{
		ID:   id,
		Name: requestBody.Name,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
