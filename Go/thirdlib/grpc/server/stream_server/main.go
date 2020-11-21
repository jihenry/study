package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "study/thirdlib/grpc/proto"
)

//StreamService ...
type StreamService struct{}

const (
	//PORT ...
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

//List ...
func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

//Record ...
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	return nil
}

//Route ...
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	return nil
}
