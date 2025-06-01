package service

import (
	"context"
	"errors"
	"log"
	"pcbook/pb"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides  Laptop service.
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

// New LaptopServer returns a new instance of LaptopServer.
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		UnimplementedLaptopServiceServer: pb.UnimplementedLaptopServiceServer{},
		Store:                            store,
	}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop reqyest with id : %s", laptop.Id)

	if len(laptop.Id) > 0 {
		//check if the laptop id is a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop Id: %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Print("client canceled the request")
		return nil, status.Error(codes.Canceled, "client canceled the request")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	//save the laptop to in-memory store

	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil

}
