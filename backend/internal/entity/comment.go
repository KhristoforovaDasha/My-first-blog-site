package entity

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	CommentId uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	PostId      uint
	UserId      uint
	CommentText string
}

type CommentRepository interface {
	Create(comment *Comment) error
	GetAll() (*[]Comment, error)
	Get(id uint) (*Comment, error)
	Update(comment *Comment) error
	Delete(id uint) error
	GetByPostId(id uint) (*[]Comment, error)
}

type CommentService interface {
	Create(comment *Comment) error
	Get(id uint) (*Comment, error)
	Update(comment *Comment) error
	Delete(comment *Comment) error
	GetCommentsByPostId(id uint) (*[]Comment, error)
}
