package repository

import (
	entity "github.com/adesokanayo/dailywellness/entity"
)

//PostRepositoryInterface contains exposed Interface
type PostRepositoryInterface interface {
	Save(post *entity.Tip) (*entity.Tip, error)
	FindAll() ([]entity.Tip, error)
	FindOne(num int64) (*entity.Tip, error)
	FindToday() (*entity.Tip, error)
}
