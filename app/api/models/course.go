package models

import (
	"capstone/app/database"
	"capstone/app/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func SaveCourse(c echo.Context, u *entity.Course) (entity.Course, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.Course{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.Course{}, errors.New("Failed to add Course")
	}

	return *u, nil
}

func FindAllCourses(c echo.Context) ([]entity.Course, error) {
	var courses []entity.Course
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&courses)
	if err.Error != nil {
		return courses, err.Error
	}
	return courses, nil
}

func FindCourseById(c echo.Context, id uint32) (entity.Course, error) {
	var u entity.Course
	db := database.GetDB(c)

	err := db.Debug().First(&u, "id = ?", id)
	if err.Error != nil {
		return entity.Course{}, errors.New("Course not Found")
	}

	return u, nil
}

func UpdateCourse(c echo.Context, id string, course *entity.Course) (int64, error) {
	db := database.GetDB(c)

	err := db.Model(&entity.Course{}).Where("id = ?", id).Updates(course)

	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func DeleteCourse(c echo.Context, id string) (int64, error) {
	db := database.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Course{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
