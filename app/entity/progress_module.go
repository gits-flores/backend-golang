package entity

import (
	"errors"
	"strings"
	"time"
	"html"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ProgressModule struct {
	gorm.Model
	EnrollCourse      EnrollCourse   `json:"enroll_course"`
	EnrollCourseID    uint32 `gorm:"not null" json:"enroll_course_id" form:"enroll_course_id"`
	Module      Module   `json:"module"`
	ModuleID    uint32 `gorm:"not null" json:"module_id" form:"module_id"`
	Status     string `gorm:"not null;" default:"Belum Selesai" json:"status" form:"status"`
}

func (u *ProgressModule) BeforeSave(c echo.Context) error {
	return nil
}

func (c *ProgressModule) Prepare() {
	c.ID = 0
	c.EnrollCourse = EnrollCourse{}
	c.Module = Module{}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.Status = html.EscapeString(strings.TrimSpace(c.Status))
}

func (u *ProgressModule) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.EnrollCourseID < 0 {
			return errors.New("Required EnrollID")
		}
		if u.ModuleID < 0 {
			return errors.New("Required ModuleID")
		}

		return nil
	}
}
