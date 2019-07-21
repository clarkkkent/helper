package main

import (
	"context"
	"errors"
	pb "github.com/clarkkkent/helper/user-service/app/proto/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Service struct {
	repo         Repository
	tokenService Authable
}

func (serv *Service) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPass)
	if err := serv.repo.Create(req); err != nil {
		return nil, err
	}
	return &pb.Response{User: req, Users:nil, Errors:nil}, nil
}

func (serv *Service) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := serv.repo.Get(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: user, Users: nil, Errors: nil}, nil
}

func (serv *Service) Auth(ctx context.Context, req *pb.User) (*pb.Token, error) {
	log.Println("Logged with:", req.Email, req.Password)
	user, err := serv.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := serv.tokenService.Encode(user)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (serv *Service) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	claims, err := serv.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		return nil, errors.New("invalid user")
	}

	return &pb.Token{Valid: true}, nil
}