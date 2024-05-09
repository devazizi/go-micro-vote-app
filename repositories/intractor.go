package repositories

import (
	"voteapp/infrastructure"

	"gorm.io/gorm"
)

type Intractor struct {
	store *gorm.DB
}

func NewDB(db infrastructure.DB) Intractor {

	return Intractor{store: db.Store}
}
