package main

import (
	"context"
	"log"
	"net"

	pb "rpc/proto"

	"google.golang.org/grpc"
)

//SearchService ...
type SearchService struct{}

//Search ...
func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

//PORT ...
const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterServiceSearchServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
