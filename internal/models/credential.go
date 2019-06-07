package models

import (
	"time"
)

type Credential struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	RegName        string `gorm:"unique"`
	HashedPassword string
}
