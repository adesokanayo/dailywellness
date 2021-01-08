package repository

import (
	entity "github.com/adesokanayo/innovation/entity"
)

//PostRepositoryInterface contains exposed Interface
type PostRepositoryInterface interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
