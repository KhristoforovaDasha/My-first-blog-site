package repository

import (
	"gorm.io/gorm"
	"hristoforovada-project/backend/internal/entity"
	repo_sqlite "hristoforovada-project/backend/internal/repository/sqlite"
	"time"
)

type Repository struct {
	User    entity.UserRepository
	Post    entity.PostRepository
	Comment entity.CommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:    repo_sqlite.NewUserSQLite(db),
		Post:    repo_sqlite.NewPostSQLite(db),
		Comment: repo_sqlite.NewCommentSQLite(db),
	}
}

func (r *Repository) DropAll(db *gorm.DB) error {
	users, err := r.User.GetAll()
	if err != nil {
		return err
	}
	for _, user := range *users {
		err = r.User.Delete(user.Id)
		if err != nil {
			return err
		}
	}

	comments, err := r.Comment.GetAll()
	if err != nil {
		return err
	}
	for _, comment := range *comments {
		err = r.Comment.Delete(comment.CommentId)
		if err != nil {
			return err
		}
	}

	posts, err := r.Post.GetAll()
	if err != nil {
		return err
	}
	for _, post := range *posts {
		err = r.Post.Delete(post.Id)
		if err != nil {
			return err
		}
	}

	db.Unscoped().Where("deleted_at > ?", time.Time{}).Delete(&entity.User{})
	db.Unscoped().Where("deleted_at > ?", time.Time{}).Delete(&entity.Comment{})
	db.Unscoped().Where("deleted_at > ?", time.Time{}).Delete(&entity.Post{})
	return nil
}
