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

func SaveModule(c echo.Context) error {
	u := new(entity.Module)

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

	moduleCreated, err := models.SaveModule(c, u)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseUser(c, utils.JSONResponseUser{
		Code:       http.StatusCreated,
		CreateUser: moduleCreated,
		Message:    "Berhasil menambahkan module",
	})
}

func Modules(c echo.Context) error {
	db := database.GetDB(c)

	Modules := []entity.Module{}
	db.Preload("Course.User").Find(&Modules)

	return c.JSON(http.StatusOK, Modules)
}

func FindModule(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB(c)

	module := entity.Module{}
	db.First(&module, id)

	return c.JSON(http.StatusOK, module)
}
