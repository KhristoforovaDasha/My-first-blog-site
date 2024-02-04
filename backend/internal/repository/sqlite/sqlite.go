package repo_sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hristoforovada-project/backend/internal/entity"
)

func NewSQLiteDB(db_uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(db_uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Post{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Comment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
