package entity

import (
	"errors"
	"time"
)

type Profile struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Photo     string    `gorm:"size:255;not null" json:"photo"`
	User      User      `json:"user"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Profile) Prepare() {
	p.ID = 0
	p.Photo = "-"
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Profile) Validate() error {

	if p.Photo == "" {
		return errors.New("Required Photo")
	}

	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}
