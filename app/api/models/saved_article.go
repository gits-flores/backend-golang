package models

import (
	"capstone/app/database"
	"capstone/app/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func SaveArticleUser(c echo.Context, u *entity.SavedArticle) (entity.SavedArticle, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.SavedArticle{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.SavedArticle{}, errors.New("Failed to add Article")
	}

	return *u, nil
}

func FindAllSavedArticles(c echo.Context) ([]entity.SavedArticle, error) {
	var article []entity.SavedArticle
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&article)
	if err.Error != nil {
		return article, err.Error
	}
	return article, nil
}
