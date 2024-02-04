package service

import (
	"hristoforovada-project/backend/internal/entity"
)

type CommentService struct {
	commentRepo entity.CommentRepository
}

func NewCommentService(commentRepo entity.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) Get(id uint) (*entity.Comment, error) {
	commentDB, err := s.commentRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return commentDB, nil
	}
}

func (s *CommentService) Delete(comment *entity.Comment) error {
	err := s.commentRepo.Delete(comment.CommentId)
	return err
}

func (s *CommentService) Create(post *entity.Comment) error {
	err := s.commentRepo.Create(post)
	return err
}

func (s *CommentService) Update(post *entity.Comment) error {
	err := s.commentRepo.Update(post)
	return err
}

func (s *CommentService) GetCommentsByPostId(id uint) (*[]entity.Comment, error) {
	commentsDB, err := s.commentRepo.GetByPostId(id)
	if err != nil {
		return nil, err
	} else {
		return commentsDB, nil
	}
}
