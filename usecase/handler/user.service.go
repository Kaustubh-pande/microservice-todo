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
	Repo         repos.Repository
	TokenService Authable
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
func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.GetAllResonse) error {
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
