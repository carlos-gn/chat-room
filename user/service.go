package user

import (
	"fmt"
)

type UserService struct {
	ur UserRepository
}

func NewService(ur UserRepository) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (s *UserService) Get(id string) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	u, err := s.ur.Get(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
