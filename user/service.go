package user

import (
	"FinalProjectGolangH8/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input domain.RegisterUserInput) (domain.User, error)
	Login(input domain.LoginUserInput) (domain.User, error)
	UpdateUser(input domain.UpdateUserInput, id int) (domain.User, error)
	GetUserByID(ID int) (domain.User, error)
	DeleteUser(id int) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterUser(input domain.RegisterUserInput) (domain.User, error) {
	user := domain.User{}
	user.Age = input.Age
	user.Email = input.Email
	user.Password = input.Password
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) Login(input domain.LoginUserInput) (domain.User, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) UpdateUser(input domain.UpdateUserInput, id int) (domain.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	user.Email = input.Email
	user.Username = input.Username
	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}
	return updateUser, nil
}

func (s *service) GetUserByID(ID int) (domain.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found on with thah ID")
	}
	return user, nil
}

func (s *service) DeleteUser(id int) (domain.User, error) {
	user, err := s.repository.DeleteUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
