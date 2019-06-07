package main

import (
	"fmt"
	"log"

	"context"

	//"time"

	pb "github.com/PandeKaustubhS/microservice-todo/usecase/user"
	"github.com/micro/go-grpc"
)

func main() {

	service := grpc.NewService()
	service.Init()

	client := pb.NewUserService("user.service", service.Client())
	clienttodo := pb.NewTodoService("user.service", service.Client())
	//id := "9"
	name := "kaustubh"
	email := "kau2345@gmail.com"
	password := "kaukau1234"

	//ts := timep.TimestampNow().
	//ts, err := timep.Timestamp(*timep.Timestamp)
	//
	//created_at := time.Now().Format("2006-01-02 15:04:05")
	//create
	r, err := client.Create(context.TODO(), &pb.User{
		//Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		//CreatedAt: ts, //created_at,
	})
	if err != nil {
		log.Fatalf("Could not create member: %v", err)
	}

	fmt.Println("Newly created user :==>", r)
	log.Println("==================================")
	r1, err := clienttodo.TodoCreate(context.TODO(), &pb.Todo{
		Task:      "test1",
		Uid:	1,
		// CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		// UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Fatalf("Could not create todo: %v", err)
	}
	fmt.Println("Newly created Todo :==>", r1)
	//get
	// get, err := client.Get(context.Background(), &pb.Getrequest{
	// 	Id: 2,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(get.User)
	log.Println("==================================")

	//getall
	log.Println("==================================")
	log.Println("==========GetAll===================")

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)

	}
	log.Println("==================================")

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Println("==================================")
	log.Printf("Your access token is: %s \n", authResponse.Token)

	fmt.Println("==================================")
	// gettodo, err := clienttodo.GetTodo(context.Background(), &pb.GetTodoRequest{
	// 	Id: 4,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get todo by id")
	// log.Println(gettodo)
	//getall
	fmt.Println("GetAll todos")
	getAlltodos, err := clienttodo.GetAllTodos(context.Background(), &pb.GetAllTodoRequest{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAlltodos.Todos {
		log.Println(v)
	}
	//update
	// fmt.Println("====================================")
	// updatetodo, err := clienttodo.UpdateTodo(context.TODO(), &pb.Todo{
	// 	Id:   3,
	// 	Task: "Test2",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("update todo by id")
	// log.Println(updatetodo)
	//Delete
	fmt.Println("====================================")
	// deletetodo, err := clienttodo.DeleteTodo(context.TODO(), &pb.Todo{
	// 	Id: 1,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("delete todo by id")
	// log.Println(deletetodo)

	//update user
	fmt.Println("====================================")
	// updateuser, err := client.Updateuser(context.TODO(), &pb.User{
	// 	Id:       1,
	// 	Name:     "kaustubh",
	// 	Password: "",
	// 	Email:    "Kaustubh@gmail.com",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("update user by id")
	// log.Println(updateuser)
	//delete user
	fmt.Println("====================================")
	// deleteuser, err := client.Deleteuser(context.TODO(), &pb.User{
	// 	Id: 2,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("delete user by id")
	// log.Println(deleteuser)
	log.Println("==================================")
	r2, err := client.GetUserTodos(context.TODO(), &pb.Getrequest{
		Id:	1,
		// CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		// UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		log.Fatalf("Could not create todo: %v", err)
	}
	fmt.Println("Related Todo :==>", r2)
}
