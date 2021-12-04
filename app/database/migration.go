package database

import (
	"capstone/app/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migration(e *echo.Echo, db *gorm.DB) {
	e.Logger.Info("Memulai dengan automigrate")

	err := db.AutoMigrate(&entity.User{})

	if err != nil {
		e.Logger.Error(err)

	}
}