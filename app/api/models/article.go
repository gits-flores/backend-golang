package models

import (
	"capstone/app/database"
	"capstone/app/entity"
	"errors"

	"github.com/labstack/echo/v4"
)

func SaveArticle(c echo.Context, u *entity.Article) (entity.Article, error) {
	db := database.GetDB(c)

	err := db.Debug().Create(&u)
	if err != nil {
		return entity.Article{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.Article{}, errors.New("Failed to add Article")
	}

	return *u, nil
}

func FindAllArticles(c echo.Context) ([]entity.Article, error) {
	var article []entity.Article
	db := database.GetDB(c)

	err := db.Debug().Limit(100).Find(&article)
	if err.Error != nil {
		return article, err.Error
	}
	return article, nil
}

func FindArticleById(c echo.Context, id uint32) (entity.Article, error) {
	var u entity.Article
	db := database.GetDB(c)

	err := db.Debug().First(&u, "id = ?", id)
	if err.Error != nil {
		return entity.Article{}, errors.New("Article not Found")
	}

	return u, nil
}

func UpdateArticle(c echo.Context, id string, article *entity.Article) (int64, error) {
	db := database.GetDB(c)

	err := db.Model(&entity.Article{}).Where("id = ?", id).Updates(article)

	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func DeleteArticle(c echo.Context, id string) (int64, error) {
	db := database.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Article{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
