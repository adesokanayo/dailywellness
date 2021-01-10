package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/adesokanayo/dailywellness/entity"
	"github.com/adesokanayo/dailywellness/repository"
)

type service struct{}

//type PostService postingInterface

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
	FindOne() (*entity.Post, error)
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
	return repo.Save(post)

}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (s *service) FindOne() (*entity.Post, error) {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 5
	num := rand.Intn(max-min+1) + min
	return repo.FindOne(int64(num))
}
