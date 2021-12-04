package utils

import "github.com/labstack/echo/v4"

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type JSONResponseUser struct {
	Code       int64       `json:"code"`
	CreateUser interface{} `json:"create_user"`
	Message    string      `json:"message"`
}

type JSONResponsesLogin struct {
	Code    int64       `json:"code"`
	Token   interface{} `json:"token"`
	Message string      `json:"message"`
}

func ResponseError(c echo.Context, err Error) error {
	return c.JSON(int(err.Code), err)
}

func ResponseUser(c echo.Context, res JSONResponseUser) error {
	return c.JSON(int(res.Code), res)
}

func ResponsesLogin(c echo.Context, res JSONResponsesLogin) error {
	return c.JSON(int(res.Code), res)
}
