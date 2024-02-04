package repo_sqlite

import (
	"errors"
	"gorm.io/gorm"
	"hristoforovada-project/backend/internal/entity"
)

type PostSQLite struct {
	db *gorm.DB
}

func NewPostSQLite(db *gorm.DB) *PostSQLite {
	return &PostSQLite{db: db}
}

func (r *PostSQLite) GetAll() (*[]entity.Post, error) {
	var posts []entity.Post

	if result := r.db.Find(&posts); result.Error != nil {
		return nil, result.Error
	} else {
		return &posts, nil
	}
}

func (r *PostSQLite) Get(id uint) (*entity.Post, error) {
	var post entity.Post

	if result := r.db.Where("id = ?", id).First(&post); result.Error == nil {
		return &post, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &post, entity.ErrPostNotFound
	} else {
		return &post, result.Error
	}
}

func (r *PostSQLite) Create(post *entity.Post) error {
	if result := r.db.Create(post); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *PostSQLite) Update(post *entity.Post) error {
	result := r.db.Model(post).Updates(post)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *PostSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Post{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
