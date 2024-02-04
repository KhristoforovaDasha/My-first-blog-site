package entity

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id         uint           `gorm:"primaryKey" json:"id"`
	Title      string         `json:"title"`
	PostText   string         `json:"post_text"`
	ImageUrl   string         `json:"imageUrl,omitempty"`
	LikesCount uint           `json:"likesCount"`
	AdminId    uint           `json:"admin_id,omitempty"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type PostRepository interface {
	Create(post *Post) error
	GetAll() (*[]Post, error)
	Get(id uint) (*Post, error)
	Update(post *Post) error
	Delete(id uint) error
}

type PostService interface {
	Create(post *Post) error
	Get(id uint) (*Post, error)
	GetAll() (*[]Post, error)
	Update(post *Post) error
	Delete(id uint) error
}
