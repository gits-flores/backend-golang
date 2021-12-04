package controllers

import (
	"capstone/app/api/models"
	"capstone/app/database"
	"capstone/app/entity"
	"capstone/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	u := new(entity.User)

	if err := c.Bind(u); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	u.Prepare()
	err := u.Validate("login")
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	token, err := SignIn(c, u.Email, u.Password)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponsesLogin(c, utils.JSONResponsesLogin{
		Code:    http.StatusCreated,
		Token:   token,
		Message: "Berhasil Login",
	})
}

func Register(c echo.Context) error {
	u := new(entity.User)

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

	err = u.BeforeSave()
	if err != nil {
		return err
	}

	userCreated, err := models.SaveUser(c, u)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseUser(c, utils.JSONResponseUser{
		Code:       http.StatusCreated,
		CreateUser: userCreated,
		Message:    "Berhasil menambahkan user",
	})
}

func SignIn(c echo.Context, email, password string) (string, error) {
	var err error
	u := new(entity.User)
	db := database.GetDB(c)

	err = db.Debug().Model(entity.User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return utils.CreateToken(c, u.ID)
}
