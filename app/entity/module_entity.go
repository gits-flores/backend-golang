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

type Module struct {
	gorm.Model
	// ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title     string `gorm:"not null;" json:"title" form:"title"`
	Thumbnail string `gorm:"not null;" json:"thumbnail" form:"thumbnail"`
	User      User   `json:"user"`
	UserID    uint32 `gorm:"not null" json:"user_id" form:"user_id"`
	Course    Course `json:"course"`
	CourseID  uint32 `gorm:"not null" json:"course_id" form:"course_id"`
}

func (m *Module) BeforeSave(context echo.Context) error {
	file, err := context.FormFile(m.Thumbnail)

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

	m.Thumbnail = string(thumb)

	return nil
}

func (m *Module) Prepare() {
	m.ID = 0
	m.Title = html.EscapeString(strings.TrimSpace(m.Title))
	m.Thumbnail = html.EscapeString(strings.TrimSpace(m.Thumbnail))
	m.User = User{}
	m.Course = Course{}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (m *Module) Validate(action string) error {
	if m.Title == "" {
		return errors.New("Required Title")
	}

	return nil
}
