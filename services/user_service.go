package services

import (
	db "gogif/db"
	repositories "gogif/repositories"
)

type UserService interface {
	Create(user *db.User) (*db.User, error)
	GetByID(id uint) (*db.User, error)
	Update(user *db.User) (*db.User, error)
	Delete(id uint) error
	List(users *[]db.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{repo: userRepo}
}

func (s *userService) Create(user *db.User) (*db.User, error) {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*db.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Update(user *db.User) (*db.User, error) {
	return s.repo.Update(user)
	}
	
func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) List(users *[]db.User) error {
	return s.repo.List(users)
}
