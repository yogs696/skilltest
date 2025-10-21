package auth

import (
	"github.com/yogs696/skilltest/internal/entity"
)

type Repository interface {
	Transaction(txFunc func(interface{}) error) (err error)

	Create(w *entity.User) (*entity.User, error)
	FindByEamil(conds map[string]interface{}) (res entity.User, row int, err error)
}
