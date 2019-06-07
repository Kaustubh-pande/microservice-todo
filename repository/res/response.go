package res

import (
	
	model "github.com/PandeKaustubhS/microservice-todo/model"
	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"

	//"bitbucket.org/rxsense/microservices-druginformation/drug-server/srv/helper"
	//"bitbucket.org/rxsense/microservices-druginformation/drug-server/srv/model"
	//pb "bitbucket.org/rxsense/microservices-druginformation/drug-server/srv/proto/drug"
	//"github.com/golang/protobuf/ptypes/wrappers"
)

func Response(resp *model.User) *pb.User {
	var user *pb.User
	
	user = &pb.User{
		// Id: resp.ID,
		Name: resp.Name,
		Email: resp.Email,
		Password: resp.Password,
		Token: resp.Token,
		// CreatedAt: resp.CreatedAt.String(),
	}
	
	return user
}

func TodoResponse(res *model.Todo)*pb.Todo{
 var todo *pb.Todo
 todo =&pb.Todo{
	 Task: res.Task,
	 Uid: res.Uid,
 }
 return todo
}

// func mapper(*model.User, )