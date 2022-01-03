package entity

import (
	"errors"
	"html"
	"io"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	// ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"not null;" json:"title" form:"title"`
	Content   string `gorm:"not null;" json:"content" form:"content"`
	Thumbnail string `gorm:"not null;" json:"thumbnail" form:"thumbnail"`
	User      User   `json:"user"`
	UserID    uint32 `gorm:"not null" json:"user_id" form:"user_id"`
}

func (u *Article) BeforeSave(c echo.Context) error {
	file, err := c.FormFile(u.Thumbnail)

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

	return nil
}

func (c *Article) Prepare() {
	c.ID = 0
	c.Title = html.EscapeString(strings.TrimSpace(c.Title))
	c.Content = html.EscapeString(strings.TrimSpace(c.Content))
	c.Thumbnail = html.EscapeString(strings.TrimSpace(c.Thumbnail))
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (u *Article) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.Title == "" {
			return errors.New("Required Title")
		}
		if u.Content == "" {
			return errors.New("Required Content")
		}

		return nil
	}
}
