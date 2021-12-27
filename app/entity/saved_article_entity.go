package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type SavedArticle struct {
	gorm.Model
	Article      Article   `json:"article"`
	ArticleID    uint32 `gorm:"not null" json:"article_id" form:"article_id"`
	User      User   `json:"user"`
	UserID    uint32 `gorm:"not null" json:"user_id" form:"user_id"`
}

func (u *SavedArticle) BeforeSave(c echo.Context) error {
	return nil
}

func (c *SavedArticle) Prepare() {
	c.ID = 0
	c.Article = Article{}
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (u *SavedArticle) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.UserID < 0 {
			return errors.New("Required UserID")
		}
		if u.ArticleID < 0 {
			return errors.New("Required ArticleID")
		}

		return nil
	}
}
