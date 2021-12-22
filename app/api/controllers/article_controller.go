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

func Articles(c echo.Context) error {
	db := database.GetDB(c)

	articles := []entity.Article{}
	db.Preload("User").Find(&articles)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)

	return c.JSON(http.StatusOK, articles)
}

func FindArticle(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB(c)

	articles := entity.Article{}
	db.Preload("User").First(&articles, id)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)

	return c.JSON(http.StatusOK, articles)
}

func SaveArticle(c echo.Context) error {
	u := new(entity.Article)

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

	articleCreated, err := models.SaveArticle(c, u)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseUser(c, utils.JSONResponseUser{
		Code:       http.StatusCreated,
		CreateUser: articleCreated,
		Message:    "Berhasil menambahkan article",
	})
}
