package app

import (
	pb "github.com/clarkkkent/helper/user-service/app/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)
const (
	port=":50051"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenservice := &TokenService{repo: repo}


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &Service{repo, tokenservice})

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
