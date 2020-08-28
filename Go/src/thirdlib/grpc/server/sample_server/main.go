package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	//1. 使用证书
	c, err := credentials.NewServerTLSFromFile("../../conf/server/server.pem", "../../conf/server/server.key")
	fpath, _ := filepath.Abs("../../conf/server.pem")
	fmt.Println(fpath)
	fmt.Println(filepath.Abs(os.Args[0]))
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(c))

	//2. 不是用证书
	// server := grpc.NewServer()

	//3. 使用CA颁发的证书
	// cert, err := tls.LoadX509KeyPair("../../conf/server/server.pem", "../../conf/server/server.key")
	// if err != nil {
	// 	log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	// }

	// certPool := x509.NewCertPool()
	// ca, err := ioutil.ReadFile("../../conf/ca.pem")
	// if err != nil {
	// 	log.Fatalf("ioutil.ReadFile err: %v", err)
	// }

	// if ok := certPool.AppendCertsFromPEM(ca); !ok {
	// 	log.Fatalf("certPool.AppendCertsFromPEM err")
	// }

	// c := credentials.NewTLS(&tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// 	ClientAuth:   tls.RequireAndVerifyClientCert,
	// 	ClientCAs:    certPool,
	// })

	// server := grpc.NewServer(grpc.Creds(c))

	pb.RegisterServiceSearchServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
