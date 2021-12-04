package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title  string `gorm:"not null;" json:"title" form:"title"`
	User   User   `json:"user"`
	UserID uint32 `gorm:"not null" json:"user_id"`
}

func (c *Course) Prepare() {
	c.ID = 0
	c.Title = html.EscapeString(strings.TrimSpace(c.Title))
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Course) Validate() error {
	if c.Title == "" {
		return errors.New("Required Title")
	}

	if c.UserID < 1 {
		return errors.New("Required Admin")
	}
	return nil
}
