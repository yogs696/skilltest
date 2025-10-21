package auth

import (
	"fmt"

	"github.com/yogs696/skilltest/internal/entity"
	"github.com/yogs696/skilltest/pkg/kemu"
	"golang.org/x/crypto/bcrypt"
)

// Service represent SKill Test auth services interface
type Service struct {
	repo Repository
	kemu *kemu.Mutex
}

// NewService creates new SKill Test auth services
func NewService(
	kemu *kemu.Mutex,
	r Repository,
	callback ...func(s string),
) *Service {
	if len(callback) > 0 {
		callback[0]("Registering User List Domain Entity...")
	}

	svc := &Service{
		repo: r,
		kemu: kemu,
	}

	return svc
}

// register user
func (s *Service) CreateUser(username, email, password string) (*entity.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	a := &entity.User{
		Username: username,
		Email:    email,
		Password: string(passwordHash),
	}

	defer func() {
		a = nil
	}()

	usr, err := s.repo.Create(a)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// Login user
func (s *Service) LoginUser(email interface{}, password string) (*entity.User, int, error) {

	conds := make(map[string]interface{})
	switch {
	case email != nil:
		conds["email"] = email

	default:
		return nil, 0, fmt.Errorf("argument email cannot be empty -> email: %d", email)
	}

	user, rows, err := s.repo.FindByEamil(conds)
	switch {
	// Always check error appear first
	case err != nil:
		return nil, rows, err

	case rows == 0:
		return nil, rows, fmt.Errorf("data with email : %s not found", email)

	default:
		return &user, rows, nil
	}
}
