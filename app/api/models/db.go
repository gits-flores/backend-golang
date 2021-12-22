package models

import (
	"capstone/app/config"
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
)

var db *sql.DB
var err error

func InitDB(e *echo.Echo) {
	config := config.GetConfig(e)

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name))
	if err != nil {
		panic("connectionString error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
