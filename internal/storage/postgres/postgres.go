package postgres

import (
	"fmt"
	url "url-shortener/internal/model"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dsn string) (*gorm.DB, error) {
	const op = "storage.postgres.New"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&url.Url{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, err
}
