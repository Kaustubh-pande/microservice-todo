package server

import (
	"github.com/PandeKaustubhS/microservice-todo/middleware"
	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
)

//
func StartGRPCServer(impl pb.UserServiceHandler, impl1 pb.TodoServiceHandler) (err error) {

	// Create a new service. Optionally include some options here.
	srv := grpc.NewService(
		micro.Name("user.service"),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
		// Middleware for error logging
		micro.WrapHandler(middleware.Logger),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterUserServiceMicroHandler(srv.Server(), impl)
	pb.RegisterTodoServiceMicroHandler(srv.Server(), impl1)
	return srv.Run()

}
