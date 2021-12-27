package utils

import (
	"capstone/app/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(c echo.Context, user_id uint32, name string) error {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() //Token expires after 1 hour

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   t,
		"data":    user_id,
		"message": "Selamat Datang " + name,
	})

}

func GetJWTData(c echo.Context, header http.Header) (jwt.MapClaims, error) {
	// var data jwtCustomClaims
	var authData string = header["Authorization"][0]
	var token string = strings.TrimPrefix(authData, "bearer ")

	finalToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfigs(c).Secret), nil
	})
	if err != nil {
		return nil, errors.New("unexpected error on getting jwt data")
	}

	claims := finalToken.Claims.(jwt.MapClaims)
	return claims, nil
}

// func GenerateToken(c echo.Context, user_id uint32) (string, error) {
// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte(config.GetConfigs(c).Secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return t, nil
// }
