package infrastructure

import (
	"context"
	"voteapp/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Store *gorm.DB
}

func InitDB(ctx context.Context, dsn string) DB {

	// isNotConnect := true

	// entities := []any{
	// 	&entity.Vote{},
	// 	&entity.VoteOption{},
	// }

	// for isNotConnect {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	mErr := db.AutoMigrate(&entity.Client{})

	if mErr != nil {
		panic(mErr)
	}

	// isNotConnect = false

	// context.WithValue(ctx, "isAvailableApp", true)

	return DB{
		Store: db,
	}

	// }
}
