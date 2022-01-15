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

type Course struct {
	gorm.Model
	// ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"not null;" json:"title" form:"title"`
	Thumbnail   string `gorm:"not null;" json:"thumbnail" form:"thumbnail"`
	Description string `gorm:"not null;" json:"description" form:"description"`
	Rangkuman string `gorm:"not null;" json:"rangkuman" form:"rangkuman"`
	User        User   `json:"user"`
	UserID      uint32 `gorm:"not null" json:"user_id" form:"user_id"`
	Modules []Module `gorm:"ForeignKey:CourseID"`
}

func (c *Course) BeforeSave(context echo.Context) error {
	file, err := context.FormFile(c.Thumbnail)

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

	c.Thumbnail = string(thumb)

	return nil
}

func (c *Course) Prepare() {
	c.ID = 0
	c.Title = html.EscapeString(strings.TrimSpace(c.Title))
	c.Thumbnail = html.EscapeString(strings.TrimSpace(c.Thumbnail))
	c.Description = html.EscapeString(strings.TrimSpace(c.Description))
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Course) Validate(action string) error {
	if c.Title == "" {
		return errors.New("Required Title")
	}

	if c.Description == "" {
		return errors.New("Required Description")
	}

	if c.Rangkuman == "" {
		return errors.New("Required Rangkuman")
	}

	return nil
}
