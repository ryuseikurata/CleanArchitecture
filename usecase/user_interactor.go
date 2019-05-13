package usecase

import "github.com/ryuseikurata/CleanArchitecture/domain"

type UserInteractor struct {
    UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

