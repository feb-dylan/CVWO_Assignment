package services

import (
	"errors"
	"forum/models"
	"forum/repositories"
	"forum/dto"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(input dto.RegisterDTO) (*dto.UserResponseDTO, error) {
	exists, err := s.userRepo.Exists(input.Username)
	if err != nil {
		return nil, errors.New("database error")
	}
	if exists {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user := &models.User{
		Username:  input.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &dto.UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *UserService) Login(input dto.LoginDTO) (*dto.UserResponseDTO, error) {
	user, err := s.userRepo.GetByUsername(input.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &dto.UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
func (s *UserService) GetUserByID(id uint) (*dto.UserResponseDTO, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &dto.UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}