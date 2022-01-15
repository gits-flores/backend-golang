package controllers

import (
	"capstone/app/api/models"
	"capstone/app/database"
	"capstone/app/entity"
	"capstone/app/utils"
	"io"
	"net/http"
	"os"
	"errors"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func SaveCourse(c echo.Context) error {
	u := new(entity.Course)

	if err := c.Bind(u); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	u.Prepare()
	err := u.Validate("")
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	file, err := c.FormFile("thumbnail")

	thumb := file.Filename
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("public/uploads/" + thumb)
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	u.Thumbnail = string(thumb)

	if err != nil {
		return err
	}

	courseCreated, err := models.SaveCourse(c, u)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseUser(c, utils.JSONResponseUser{
		Code:       http.StatusCreated,
		CreateUser: courseCreated,
		Message:    "Berhasil menambahkan course",
	})
}

func Courses(c echo.Context) error {
	db := database.GetDB(c)

	Courses := []entity.Course{}
	db.Preload("Modules").Preload("User").Find(&Courses)

	return c.JSON(http.StatusOK, Courses)
}

func FindCourse(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB(c)

	course := entity.Course{}
	db.Preload("Modules").Preload("User").First(&course, id)

	return c.JSON(http.StatusOK, course)
}

func EnrollCourse(c echo.Context) error {
	u := new(entity.EnrollCourse)

	if err := c.Bind(u); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	u.Prepare()
	err := u.Validate("")
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	db := database.GetDB(c)
	enrol := entity.EnrollCourse{}

	result := db.Where("user_id = ?", c.FormValue("user_id")).Where("course_id = ?", c.FormValue("course_id")).Take(&enrol)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.ResponseUser(c, utils.JSONResponseUser{
			Code:       http.StatusCreated,
			CreateUser: "Course sudah diikuti",
			Message:    "Course sudah diikuti",
		})
	} else {
		enroll, err := models.EnrollCourseUser(c, u)

		if err != nil {
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	
		return utils.ResponseUser(c, utils.JSONResponseUser{
			Code:       http.StatusCreated,
			CreateUser: enroll,
			Message:    "Berhasil mengikuti course",
		})
	}
}