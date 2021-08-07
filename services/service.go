package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/adesokanayo/dailywellness/entity"
	"github.com/adesokanayo/dailywellness/repository"
)

type service struct{}

var (
	repo = repository.NewFireStoreRepository()
)

func NewTipService() postingInterface {
	return &service{}
}

type postingInterface interface {
	Validate(post *entity.Tip) error
	Create(post *entity.Tip) (*entity.Tip, error)
	FindAll() ([]entity.Tip, error)
	FindOne() (*entity.Tip, error)
	FindToday() (*entity.Tip, error)
}

func (s *service) Validate(post *entity.Tip) error {

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
func (s *service) Create(post *entity.Tip) (*entity.Tip, error) {

	post.ID = rand.Int63()
	return repo.Save(post)

}

func (s *service) FindAll() ([]entity.Tip, error) {
	return repo.FindAll()
}

func (s *service) FindOne() (*entity.Tip, error) {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	num := rand.Intn(max-min+1) + min
	return repo.FindOne(int64(num))
}

func (s *service) FindToday() (*entity.Tip, error) {

	return repo.FindToday()
}
