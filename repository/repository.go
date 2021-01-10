package repository

import (
	entity "github.com/adesokanayo/dailywellness/entity"
)

//PostRepositoryInterface contains exposed Interface
type PostRepositoryInterface interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindOne(num int64 )(*entity.Post, error)
}
