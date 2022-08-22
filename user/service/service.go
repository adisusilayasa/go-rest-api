package service

import (
	"belajar-interface/user/model"
	"belajar-interface/user/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type M map[string]interface{}

type Service interface {
	AuthenticateUser(username, password string) (bool, model.LeakUserData)
	CheckPasswordHash(password, hash string) bool
	CreateUser(user model.UserData) (model.UserData, error)
	PasswordHashing(password string) (string, error)
}
type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) AuthenticateUser(username, password string) (bool, model.LeakUserData) {
	user, err := s.repository.GetUser(username)
	if err != nil {
		panic(err)
	}
	var UserDataResponse model.LeakUserData
	if user.Username == username {
		res := s.CheckPasswordHash(user.Password, password)
		if res == true {
			UserDataResponse.UID = user.UID
			UserDataResponse.Role = user.Role
			UserDataResponse.Username = user.Username
			UserDataResponse.Email = user.Email
			s.repository.UpdateLastLogin(user.UID)
			return true, UserDataResponse
		}
	}

	return false, UserDataResponse
}

func (s *service) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *service) CreateUser(user model.UserData) (model.UserData, error) {
	newUser := model.UserData{}
	password, err := s.PasswordHashing(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	newUser.UID = user.UID
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.Password = password

	res, err := s.repository.CreateUser(newUser)
	log.Print(newUser)
	if err != nil {
		panic(err)
	}
	return res, nil
}
func (s *service) PasswordHashing(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}
