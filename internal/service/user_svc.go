package service

import (
	"errors"

	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(data *models.CreateUserRequest) (*models.User, error) {

	if data.Email == "" ||
		data.Password == "" ||
		data.Username == "" ||
		data.Phone == "" {
		return nil, errors.New("all fields are required")
	}

	if len(data.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	if s.repo.FindByEmail(data.Email) != nil {
		return nil, errors.New("email already exists")
	}

	user := s.repo.Create(data)

	return user, nil
}

func (s *UserService) Login(users *models.LoginRequest) (*models.User, error) {
	user := s.repo.FindByEmail(users.Email)

	if user == nil {
		return nil, errors.New("email not found")
	}

	if user.Password != users.Password {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
