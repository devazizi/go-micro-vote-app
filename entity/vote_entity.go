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
	Name   string
	VoteId int
	Vote   Vote
}

type ClientVote struct {
	gorm.Model
	ClientId     int
	VoteId       int
	VoteOptionId int
	Description  string
	Client       Client
	Vote         Vote
	VoteOption   VoteOption
}

type Client struct {
	gorm.Model
	CellNumber   string `gorm:"uniqueIndex"`
	PhoneNumber  string `gorm:"uniqueIndex"`
	NationalCode string `gorm:"uniqueIndex"`
	EmailAddress string `gorm:"uniqueIndex"`
	Password     string
}
