package service

import (
	"hristoforovada-project/backend/internal/entity"
)

type PostService struct {
	postRepo entity.PostRepository
}

func NewPostService(postRepo entity.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) Get(id uint) (*entity.Post, error) {
	postDB, err := s.postRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return postDB, nil
	}
}

func (s *PostService) GetAll() (*[]entity.Post, error) {
	postsDB, err := s.postRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		return postsDB, nil
	}
}

func (s *PostService) Delete(id uint) error {
	err := s.postRepo.Delete(id)
	return err
}

func (s *PostService) Create(post *entity.Post) error {
	err := s.postRepo.Create(post)
	return err
}

func (s *PostService) Update(post *entity.Post) error {
	err := s.postRepo.Update(post)
	return err
}
