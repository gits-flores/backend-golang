package database

import (
	"capstone/app/config"

	"fmt"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func Init(e *echo.Echo) {

	e.Logger.Info("menginisialisasikan database")

	config := config.GetConfig(e)
	e.Logger.Info(config)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		e.Logger.Error(err)

	}
	Migration(e, db)
}

func GetDB(c echo.Context) *gorm.DB {
	if db == nil {
		c.Logger().Error("db belum terinisilisasi")
	}
	return db
}
