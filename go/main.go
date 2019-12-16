package main

import (
	"context"
	"fmt"
	// "github.com/appoptics/appoptics-apm-go/v1/ao"
	pb "github.com/okapusta/grpc-workshop/go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type myAwesomeServer struct {
	pb.UnimplementedMyAwesomeServiceServer
}

func (s *myAwesomeServer) GetSomeData(ctx context.Context, req *pb.GetSomeDataRequest) (*pb.GetSomeDataResponse, error) {
	id := req.GetId()

	fmt.Printf("Requested ID: %d\n", id)

	// let's build response
	tags := pb.Tags{
		Tags: map[string]string{
			"job": "developer",
			"likes": "pizza",
		},
	}
	response := pb.GetSomeDataResponse{
		Id: id,
		Name:  "Oskar",
		Tags: &tags,
	}
	return &response, nil
}

func main() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Start\n")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := myAwesomeServer{}
	pb.RegisterMyAwesomeServiceServer(grpcServer, &s)
	grpcServer.Serve(lis)
}
