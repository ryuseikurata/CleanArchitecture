package usecase

import "github.com/ryuseikurata/CleanArchitecture/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}