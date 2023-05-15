package usecase

import (
	"docker-echo-template/api/domain"
)

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}
