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
	Updateuser(user *pb.User) error
	Deleteuser(*pb.User) error
	GetByEmail(email string) (*pb.User, error)
}
type TodoRepository interface {
	TodoCreate(todo *pb.Todo) error
	GetTodo(id int32) (*pb.Todo, error)
	GetAllTodo() ([]*pb.Todo, error)
	UpdateTodo(todo *pb.Todo) error
	DeleteTodo(todo *pb.Todo) error
}

//
type UserRepository struct {
	Db *gorm.DB
}

//create user
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err

	}
	return nil
}

//getuser
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

//getall users
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//getbyemail
func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.Db.Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//update user
func (repo *UserRepository) Updateuser(user *pb.User) error {
	if err := repo.Db.Model(&user).Updates(user).Error; err != nil {
		return err

	}
	return nil
}

//delete user
func (repo *UserRepository) Deleteuser(user *pb.User) error {
	if err := repo.Db.Delete(&user).Error; err != nil {
		return err

	}
	return nil
}

//Todo create
func (repo *UserRepository) TodoCreate(todo *pb.Todo) error {

	if err := repo.Db.Create(todo).Error; err != nil {
		return err

	}
	return nil
}
func (repo *UserRepository) GetTodo(id int32) (*pb.Todo, error) {
	repo.Db.LogMode(true)
	todo := &pb.Todo{}
	todo.Id = id
	if err := repo.Db.Find(&todo).Error; err != nil {
		return nil, err
	}
	repo.Db.LogMode(true)
	return todo, nil
}
func (repo *UserRepository) GetAllTodo() ([]*pb.Todo, error) {
	var todos []*pb.Todo
	if err := repo.Db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
func (repo *UserRepository) UpdateTodo(todo *pb.Todo) error {
	if err := repo.Db.Model(&todo).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) DeleteTodo(todo *pb.Todo) error {
	if err := repo.Db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}
