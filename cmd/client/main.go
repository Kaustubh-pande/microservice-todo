package main

import (
	"fmt"
	"log"

	"context"

	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"github.com/micro/go-grpc"
)

func main() {

	service := grpc.NewService()
	service.Init()

	client := pb.NewUserService("user.service", service.Client())
	//id := "9"
	name := "kaustubh"
	email := "kau2345@gmail.com"
	password := "kaukau1234"
	//create
	r, err := client.Create(context.TODO(), &pb.User{
		//Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not create member: %v", err)
	}
	fmt.Println(r)
	log.Println("==================================")

	//get
	// get, err := client.Get(context.Background(), &pb.Getrequest{
	// 	Id: 2,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(get.User)
	// //getall
	// getAll, err := client.GetAll(context.Background(), &pb.Request{})
	// if err != nil {
	// 	log.Fatalf("Could not list users: %v", err)
	// }
	// for _, v := range getAll.Users {
	// 	log.Println(v)

	// }

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Println("==================================")
	log.Printf("Your access token is: %s \n", authResponse.Token)

}
