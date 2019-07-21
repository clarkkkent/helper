package main

import (
	"context"
	"flag"
	pb "github.com/clarkkkent/helper/user-service/app/proto/user"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err!=nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	namePtr := flag.String("username", "", "name for user")
	emailPtr := flag.String("email", "", "email for user")
	passwordPtr := flag.String("password", "", "password for user")
	flag.Parse()

	log.Println(*namePtr, *emailPtr, *passwordPtr)

	_, err = client.Create(context.Background(), &pb.User{
		Name: *namePtr,
		Email: *emailPtr,
		Password: *passwordPtr,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
}

