package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type EnrollCourse struct {
	gorm.Model
	ID      uint32   `json:"id"`
	Course      Course   `json:"course"`
	CourseID    uint32 `gorm:"not null" json:"course_id" form:"course_id"`
	User      User   `json:"user"`
	UserID    uint32 `gorm:"not null" json:"user_id" form:"user_id"`
}

func (u *EnrollCourse) BeforeSave(c echo.Context) error {
	return nil
}

func (c *EnrollCourse) Prepare() {
	c.ID = 0
	c.Course = Course{}
	c.User = User{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (u *EnrollCourse) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.UserID < 0 {
			return errors.New("Required UserID")
		}
		if u.CourseID < 0 {
			return errors.New("Required CourseID")
		}

		return nil
	}
}
