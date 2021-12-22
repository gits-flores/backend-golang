package routes

import (
	"capstone/app/api/controllers"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) *echo.Echo {

	auth := e.Group("")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	auth.GET("/articles", controllers.Articles)
	auth.GET("/articles/:id", controllers.FindArticle)
	auth.POST("/article", controllers.SaveArticle)

	return e
}
