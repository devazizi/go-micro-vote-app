package infrastructure

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Store *gorm.DB
}

func InitDB(ctx context.Context, dsn string) DB {

	isNotConnect := true

	// entities := []any{
	// 	&entity.Vote{},
	// 	&entity.VoteOption{},
	// }

	for isNotConnect {

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			// context
		}

		// mErr := db.AutoMigrate(...entities)

		// if mErr != nil {

		// }

		isNotConnect = false

		context.WithValue(ctx, "isAvailableApp", true)

		return DB{
			Store: db,
		}

	}
}
