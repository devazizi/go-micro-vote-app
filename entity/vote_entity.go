package entity

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	Name        string
	Description string
	EndTime     time.Time
}

type VoteOption struct {
	gorm.Model
	Name string
	Vote
}
