package categories

import (
	c_repository "github.com/GabrielBrotas/go-categories-msvc/internal/categories/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(r *gin.Engine, db *gorm.DB) {
	repository := c_repository.NewCategoryRepository(db)

	r.GET("/categories", func(c *gin.Context) {
		getCategories(c, repository)
	})

	r.GET("/categories/:id", func(c *gin.Context) {
		getCategory(c, repository)
	})

	r.POST("/categories", func(c *gin.Context) {
		createCategory(c, repository)
	})

	r.PATCH("/categories/:id", func(c *gin.Context) {
		updateCategory(c, repository)
	})

	r.DELETE("/categories/:id", func(c *gin.Context) {
		deleteCategory(c, repository)
	})
}
