package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"yuthi.com/mailService/proto"
	"yuthi.com/mailService/services"
)


func main() {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5000")
	proto.RegisterMailServiceServer(server, services.AuthServer{})
	if err != nil {
		log.Fatal("Error Creating listener", err.Error())
	}
	server.Serve(listener)
}