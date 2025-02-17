package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(dto RegisterUserDto) (User, error)
	LoginUser(dto LoginUserDto) (User, error)
	isEmailAvailable(dto CheckEmailDto) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	SaveToken(ID int, token string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}

}

func (s *service) RegisterUser(dto RegisterUserDto) (User, error) {
	user := User{}
	user.Name = dto.Name
	user.Occupation = dto.Ocupation
	user.Email = dto.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = "user"

	// Check if email already exists
	checkUser, err := s.repository.FindByEmail(dto.Email)

	if err != nil {
		return checkUser, err
	}

	if len(checkUser.Email) > 0 {
		return checkUser, errors.New("Email already registered")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(dto LoginUserDto) (User, error) {
	user, err := s.repository.FindByEmail(dto.Email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found")

	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) isEmailAvailable(dto CheckEmailDto) (bool, error) {

	user, err := s.repository.FindByEmail(dto.Email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	user.Avatar = fileLocation

	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) SaveToken(ID int, token string) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	user.Token = token

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}
