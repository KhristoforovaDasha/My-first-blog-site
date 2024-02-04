package service

import (
	"hristoforovada-project/backend/internal/entity"
	"hristoforovada-project/backend/internal/repository"
)

type Service struct {
	User    entity.UserService
	Post    entity.PostService
	Comment entity.CommentService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Post:    NewPostService(repo.Post),
		Comment: NewCommentService(repo.Comment),
	}
}
