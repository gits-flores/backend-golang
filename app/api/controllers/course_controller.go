package controllers

import (
	"capstone/app/api/models"
	"capstone/app/database"
	"capstone/app/entity"
	"capstone/app/utils"
	"io"
	"net/http"
	"os"

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
	db.Preload("User").Find(&Courses)

	return c.JSON(http.StatusOK, Courses)
}

func FindCourse(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB(c)

	course := entity.Course{}
	db.Preload("User").First(&course, id)

	return c.JSON(http.StatusOK, course)
}
