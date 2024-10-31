// FLOW: ya ini bikin kaya biasa wkwk, ini jadi blueprint
package services

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound       = errors.New("data not found")
	ErrInternalServer = errors.New("internal server error")
)

type User struct {
	Id   int
	Name string
}

type Repository interface {
	GetAll() ([]User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetAll() (users []User, err error) {
	users, err = s.repo.GetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		} else {
			return nil, ErrInternalServer
		}
	}

	if len(users) == 0 {
		return nil, errors.New("data not found")
	}
	return
}
