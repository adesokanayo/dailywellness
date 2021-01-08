package services

import (
	"errors"
	"math/rand"

	"github.com/adesokanayo/innovation/entity"
	"github.com/adesokanayo/innovation/repository"
)

type service struct{}
type PostService postingInterface

var (
	repo = repository.NewFireStoreRepository()
)

func NewPostService() postingInterface {
	return &service{}
}

type postingInterface interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func (s *service) Validate(post *entity.Post) error {

	if post == nil {
		err := errors.New("empty post")
		return err
	}

	if post.Title == "" {
		err := errors.New("empty title")
		return err
	}

	return nil
}
func (s *service) Create(post *entity.Post) (*entity.Post, error) {

	post.ID = rand.Int63()
	return repository.FireStoreRepo.Save(post)

}

func (s *service) FindAll() ([]entity.Post, error) {
	return repository.FireStoreRepo.FindAll()
}
