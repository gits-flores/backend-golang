package models

import (
	"capstone/app/database"
	"capstone/app/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func EnrollCourseUser(c echo.Context, u *entity.EnrollCourse) (entity.EnrollCourse, error) {
	db := database.GetDB(c)

	

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.EnrollCourse{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.EnrollCourse{}, errors.New("Failed to add Article")
	}

	return *u, nil
}

func FindAllEnrollCourses(c echo.Context) ([]entity.EnrollCourse, error) {
	var course []entity.EnrollCourse
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&course)
	if err.Error != nil {
		return course, err.Error
	}
	return course, nil
}
