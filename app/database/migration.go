package database

import (
	"capstone/app/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migration(e *echo.Echo, db *gorm.DB) {
	e.Logger.Info("Memulai dengan automigrate")

	err := db.AutoMigrate(&entity.User{}, &entity.Article{}, &entity.SavedArticle{}, &entity.Course{}, &entity.Module{}, &entity.EnrollCourse{}, &entity.ProgressModule{})

	if err != nil {
		e.Logger.Error(err)

	}
}
