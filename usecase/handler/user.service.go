package usecasehadler

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	repos "github.com/PandeKaustubhS/microservice-todo/repository"
	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

//service
type Service struct {
	Repo repos.Repository

	TokenService Authable
}
type TodoService struct {
	TodoRepo repos.TodoRepository
}

func (srv *Service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	log.Println("In create")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}
	req.Password = string(hashedPass)
	if err := srv.Repo.Create(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.TokenService.Encode(req)
	if err != nil {
		return err
	}

	res.User = req
	res.Token = &pb.Token{Token: token}

	// if err := srv.Repo.Create(req); err != nil {
	// 	return nil
	// }
	/*
		if err := srv.Publisher.Publish(ctx, req); err != nil {
			return errors.New(fmt.Sprintf("error publishing event: %v", err))
		}*/

	return nil

}

//get
func (srv *Service) Get(ctx context.Context, req *pb.Getrequest, res *pb.GetResponse) error {
	fmt.Println(req.Id)
	user, err := srv.Repo.Get(req.Id)

	if err != nil {
		return err
	}
	fmt.Println(user)
	res.User = user
	return nil
}

//getall
func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.GetAllResponse) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *Service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("logging with: ", req.Email, req.Password)
	user, err := srv.Repo.GetByEmail(req.Email)
	log.Println(user, err)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.TokenService.Encode(user)
	if err != nil {
		return err
	}
	fmt.Println(token)
	res.Token = token //&pb.Token{Token: token}
	return nil
}

//
func (srv *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.TokenService.Decode(req.Token)

	if err != nil {
		return err
	}
	str := claims.User.Id
	if strconv.Itoa(int(str)) == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}

//update user
func (srv *Service) Updateuser(ctx context.Context, req *pb.User, res *pb.Response) error {
	log.Println("In update")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}
	req.Password = string(hashedPass)
	if err := srv.Repo.Updateuser(req); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := srv.TokenService.Encode(req)
	if err != nil {
		return err
	}

	res.User = req
	res.Token = &pb.Token{Token: token}

	return nil
}

//delete user
func (srv *Service) Deleteuser(ctx context.Context, req *pb.User, res *pb.DeleteResponse) error {
	log.Println("Deleting todo")
	err := srv.Repo.Deleteuser(req)
	if err != nil {
		return err
	}

	res.Success = true
	res.User = req
	return nil
}

//todo
func (srv *TodoService) TodoCreate(ctx context.Context, req *pb.Todo, res *pb.TodoResponse) error {
	log.Println("Creating todo")
	err := srv.TodoRepo.TodoCreate(req)
	if err != nil {
		return err
	}

	res.Todo = req
	return nil
}

//get todo
func (srv *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest, res *pb.GetTodoResponse) error {
	fmt.Println(req.Id)
	todo, err := srv.TodoRepo.GetTodo(req.Id)

	if err != nil {
		return err
	}
	fmt.Println(todo)
	res.Todo = todo
	return nil
}

//getall todo
func (srv *TodoService) GetAllTodos(ctx context.Context, req *pb.GetAllTodoRequest, res *pb.GetAllTodoResponse) error {
	todos, err := srv.TodoRepo.GetAllTodo()
	if err != nil {
		return err
	}
	res.Todos = todos
	return nil
}

//update todo
func (srv *TodoService) UpdateTodo(ctx context.Context, req *pb.Todo, res *pb.TodoResponse) error {
	log.Println("Creating todo")
	err := srv.TodoRepo.UpdateTodo(req)
	if err != nil {
		return err
	}

	res.Todo = req
	return nil
}

//delete todo

func (srv *TodoService) DeleteTodo(ctx context.Context, req *pb.Todo, res *pb.DeleteTodoResponse) error {
	log.Println("Deleting todo")
	err := srv.TodoRepo.DeleteTodo(req)
	if err != nil {
		return err
	}

	res.Message = true
	return nil
}
