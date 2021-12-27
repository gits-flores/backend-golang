package routes

import (
	"capstone/app/api/controllers"
	"capstone/app/middleware"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) *echo.Echo {

	auth := e.Group("", middleware.IsAuthenticated)

	auth.GET("/articles", controllers.Articles)
	auth.GET("/articles/:id", controllers.FindArticle)
	auth.POST("/article", controllers.SaveArticle)

	auth.POST("/save-article-user", controllers.SaveArticleUser)
	auth.GET("/save-article-user/:id", controllers.GetArticleUser)
	auth.GET("/cek-article-user/:user_id/:article_id", controllers.CekArticleUser)
	auth.GET("/delete-article-user/:user_id/:article_id", controllers.DeleteArticleUser)

	return e
}
