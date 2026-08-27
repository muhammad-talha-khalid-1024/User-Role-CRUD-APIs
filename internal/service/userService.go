package service

import (
	secretPassword "mustaqel/pkg/utils"

	"mustaqel/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Check if user already exists
	existing, _ := s.repo.GetByEmail(req.Email)
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := secretPassword.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      req.Age,
		Status:   req.Status,
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	// Don't return password
	user.Password = ""
	return user, nil
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(id string, req *domain.UpdateUserRequest) (*domain.User, error) {
	// Get existing user
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Email != "" {
		existing.Email = req.Email
	}
	if req.Age > 0 {
		existing.Age = req.Age
	}
	if req.Status >= 0 {
		existing.Status = req.Status
	}

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) ListUsers(page, pageSize int) ([]*domain.User, int, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	users, total, err := s.repo.List(pageSize, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (total + pageSize - 1) / pageSize
	return users, total, totalPages, nil
}
