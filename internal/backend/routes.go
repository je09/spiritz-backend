package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/je09/spritz-backend/internal/backend/apiV1"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/info", apiV1.GetBookInfo)

		v1.Group("/library")
		{
			v1.GET("/categories", apiV1.GetCategories)
			v1.GET("/categories/popular", apiV1.GetBooksByCategory)

			v1.POST("/bookparse/borrow", apiV1.BorrowPublicBook)
			v1.GET("/bookparse/info", apiV1.GetBookInfo)
			v1.GET("/bookparse/text", apiV1.GetText)
			v1.POST("/bookparse", apiV1.CreateNewBook)
			v1.POST("/bookparse/text", apiV1.CreateNewText)
			v1.DELETE("/bookparse", apiV1.DeleteBook)
		}

		v1.Group("/progress")
		{
			v1.GET("/", apiV1.GetAvgStat)
			v1.GET("/bookparse", apiV1.GetStatByBook)
			v1.PUT("/bookparse", apiV1.CreateOrUpdateBookStat)
		}
	}

	return router
}
