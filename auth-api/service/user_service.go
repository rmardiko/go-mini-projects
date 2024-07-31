package service

import (
	"errors"
	"log"
	"rmardiko/auth-api/models"
	repos "rmardiko/auth-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repos.UserRepository
}

func NewUserService(u repos.UserRepository) UserService {
	return UserService{
		UserRepo: u,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {

	users, err := s.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) AddUser(newUser models.User) error {

	encryptedPass, err := hashAndSalt([]byte(newUser.Password))
	if err != nil {
		return err
	}

	// replace the password with the encrypted
	newUser.Password = encryptedPass

	err = s.UserRepo.Insert(newUser)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ValidateUser(userParam models.User) error {

	userData, err := s.UserRepo.FindByName(userParam.Username)

	if err != nil {
		return err
	} else {
		if comparePasswords(userData.Password, []byte(userParam.Password)) {
			return nil
		}

		return errors.New("cannot verify credentials")
	}
}

func hashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
