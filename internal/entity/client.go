package entity

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	Id    uint64 `gorm:"primaryKey"`
	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
