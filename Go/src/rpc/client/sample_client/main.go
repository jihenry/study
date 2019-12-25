package main

import (
	"context"
	"log"

	pb "rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//PORT ...
const PORT = "9001"

func main() {
	//1. 使用证书，第二个参数生成证书指定的Name
	c, err := credentials.NewClientTLSFromFile("../../conf/server/server.pem", "go-grpc-viking")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	//2. 不使用证书
	// conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	//3. 使用CA颁发的证书
	// cert, err := tls.LoadX509KeyPair("../../conf/client/client.pem", "../../conf/client/client.key")
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
	// 	ServerName:   "go-grpc-viking",
	// 	RootCAs:      certPool,
	// })
	// conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewServiceSearchClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
