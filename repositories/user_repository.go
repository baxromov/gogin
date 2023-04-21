package repositories

import (
	"fmt"

	"gorm.io/gorm"
    db "gogif/db"
)

type UserRepository interface {
	Create(user *db.User) (*db.User, error)
	GetByID(id uint) (*db.User, error)
	Update(user *db.User) (*db.User, error)
	Delete(id uint) error
	List(users *[]db.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{db: database}
}

func (repo *userRepository) Create(user *db.User) (*db.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func (repo *userRepository) GetByID(id uint) (*db.User, error) {
	var user db.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &user, nil
}

func (repo *userRepository) Update(user *db.User) (*db.User, error) {
	if err := repo.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return user, nil
}

func (repo *userRepository) Delete(id uint) error {
	if err := repo.db.Delete(&db.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

func (repo *userRepository) List(users *[]db.User) error {
	if err := repo.db.Find(users).Error; err != nil {
		return fmt.Errorf("failed to list users: %v", err)
	}
	return nil
}
