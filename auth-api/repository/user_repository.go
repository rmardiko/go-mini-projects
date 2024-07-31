package repository

import (
	"errors"
	"rmardiko/auth-api/models"
)

type UserRepository interface {
	Insert(user models.User) error
	FindById(id int32) (models.User, error)
	FindAll() ([]models.User, error)
	FindByName(userName string) (models.User, error)
}

type SliceUserRepo struct {
	users []models.User
}

func NewSliceUserRepository() UserRepository {
	return &SliceUserRepo{
		users: []models.User{},
	}
}

func (repo *SliceUserRepo) Insert(newUser models.User) error {
	repo.users = append(repo.users, newUser)

	return nil
}

func (repo *SliceUserRepo) FindAll() ([]models.User, error) {
	return repo.users, nil
}

func (repo *SliceUserRepo) FindByName(username string) (models.User, error) {

	for _, u := range repo.users {

		if u.Username == username {
			return u, nil
		}
	}

	// return empty
	return models.User{}, errors.New("record not found")
}

func (repo *SliceUserRepo) FindById(id int32) (models.User, error) {
	for _, u := range repo.users {

		if u.ID == string(id) {
			return u, nil
		}
	}

	// return empty
	return models.User{}, errors.New("record not found")
}
