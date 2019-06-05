GOPATH:=$(shell printenv GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
PROTOCMD=protoc

grpc:
	# building grpc binary
	$(GOBUILD) -o bin/grpc cmd/grpc/main.go
rest:
	# building rest gateway binary
	$(GOBUILD) -o bin/rest cmd/rest/main.go
#github.com/PandeKaustubhS/
proto:
	# compiling protobuffer code:
	protoc --proto_path=api \
	-I /home/synerzip/Documents/microservice-todo/api \
	--micro_out=./usecase/user \
	--go_out=plugins=grpc:./usecase/user usertodo.proto

	#compiling rest gateway
	protoc --proto_path=api \
	-I /home/synerzip/Documents/microservice-todo/api \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=api/user.yaml:./usecase/user usertodo.proto

	# # compiling open api info
	protoc --proto_path=api \
	-I /home/synerzip/Documents/microservice-todo/api \
	--swagger_out=logtostderr=true:api usertodo.proto

	sed -i -e "s/RegisterUserServiceHandler/RegisterUserServiceMicroHandler/g" ./usecase/user/usertodo.micro.go

	sed -i -e "s/RegisterTodoServiceHandler/RegisterTodoServiceMicroHandler/g" ./usecase/user/usertodo.micro.go