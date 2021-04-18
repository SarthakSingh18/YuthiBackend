package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	GreeterService "github.com/myuser/myrepo/proto/helloworld"
	HttpBodyExampleService "github.com/myuser/myrepo/proto/helloworld"
	createOrganisationService "github.com/myuser/myrepo/proto/helloworld"
	services "github.com/myuser/myrepo/services"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"net/http"
)

type server struct{
	GreeterService.UnimplementedGreeterServer
}
type server1 struct{
	HttpBodyExampleService.UnimplementedHttpBodyExampleServiceServer
}

func (s1 *server1) HelloWorld(ctx context.Context, e *emptypb.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:  []byte("Hello World"),
	}, nil
}


func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *GreeterService.HelloRequest) (*GreeterService.HelloReply, error) {
	log.Println("I am here")
	log.Println(in.Name)
	return &GreeterService.HelloReply{Message: in.Name + " world"}, nil
}


func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	GreeterService.RegisterGreeterServer(s, &server{})
	HttpBodyExampleService.RegisterHttpBodyExampleServiceServer(s, &server1{})
	createOrganisationService.RegisterOrganisationServer(s,&services.Server2{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = GreeterService.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}