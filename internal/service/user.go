package service

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) SignUp(user *model.User) bool {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "sign up",
		}).Error("Failed to hash password")
		return false
	}

	account := model.Credential{
		RegisterName:   user.RegisterName,
		HashedPassword: string(hashedPassword),
	}

	ok := s.userRepo.Create(user, &account)
	if !ok {
		return false
	}
	return true
}

func (s *UserService) LogIn(user *model.User) bool {
	_, credential, ok := s.userRepo.FindByRegisterName(user.RegisterName)
	if !ok {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(user.Password)); err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "log in",
		}).Error("Wrong password")
		return false
	}
	return true
}
