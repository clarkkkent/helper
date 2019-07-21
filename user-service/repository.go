package main

import (
	pb "github.com/clarkkkent/helper/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(*pb.User) error
	Get(string) (*pb.User, error)
	GetByEmail(string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func(repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err!=nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
