package repository

import (
	// "github.com/gogo/protobuf/proto"
	model "github.com/PandeKaustubhS/microservice-todo/model"
	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"github.com/jinzhu/gorm"
	res "github.com/PandeKaustubhS/microservice-todo/repository/res"
	"log"
)

//
type Repository interface {
	Get(id int32) (*pb.User, error)
	GetUserTodos(id int32) ([]*pb.Todo, error)
	Create(user *pb.User)(*pb.User,error)
	GetAll() ([]*pb.User, error)
	Updateuser(user *pb.User) error
	Deleteuser(*pb.User) error
	GetByEmail(email string) (*pb.User, error)
}
type TodoRepository interface {
	TodoCreate(todo *pb.Todo) (*pb.Todo,error)
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
func (repo *UserRepository) Create(user *pb.User) (*pb.User,error) {
	u:= model.User{Name:user.Name, Email:user.Email, Password:user.Password, Token:user.Token}
	// log.Println("1")

	// log.Println(readJSON(user, u))
	
	if err := repo.Db.Create(&u).Error; err != nil {
		return nil,err
	}
	// log.Println("2")
	u1 := res.Response(&u) 
	return u1,nil
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
	var users []*model.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	response := []*pb.User{}
	for _, v := range users {
		response = append(response,res.Response(v))
		}	
		return response, nil
}

//getbyemail
func (repo *UserRepository) GetByEmail(Email string) (*pb.User, error) {
	user := &model.User{}
	if err := repo.Db.Where("email = ?", Email).
		First(&user).Error; err != nil {
		return nil, err
	}
	res := res.Response(user)
	return res, nil
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
func (repo *UserRepository) TodoCreate(todo *pb.Todo) (*pb.Todo,error) {
	t:= model.Todo{Task:todo.Task,Uid:todo.Uid}
	if err := repo.Db.Create(&t).Error; err != nil {
		return nil,err

	}
	res := res.TodoResponse(&t)
	return res,nil
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
func (repo *UserRepository) GetUserTodos(id int32) ([]*pb.Todo, error) {
	user := &model.User{}
	var todos []*model.Todo
	if err := repo.Db.Table("users").Where("users.id = ?", id).
	First(&user).Related(&todos).Error; err != nil {
		return nil, err
	}
	log.Println("Todos", todos)
	repo.Db.LogMode(true)
	// return user, nil
	return nil, nil
}