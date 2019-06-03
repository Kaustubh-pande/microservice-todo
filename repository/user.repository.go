package repository

import (
	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"github.com/jinzhu/gorm"
)

//
type Repository interface {
	Get(id int32) (*pb.User, error)
	Create(user *pb.User) error
	GetAll() ([]*pb.User, error)
	//update(user *pb.User) error
	//delete(id string) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
}

//
type UserRepository struct {
	Db *gorm.DB
}

//
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err

	}
	return nil
}
func (repo *UserRepository) Get(id int32) (*pb.User, error) {
	repo.Db.LogMode(true)
	user := &pb.User{}
	user.Id = id
	if err := repo.Db.Find(&user).Error; err != nil {
		return nil, err
	}
	repo.Db.LogMode(true)
	return user, nil
}
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.Db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
