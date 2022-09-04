package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc_micro/b_micro/b_protopb"
	"grpc_micro/b_micro/server"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8080", "port number")
	flag.Parse()

	log.Println("Starting server on port " + *port)
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	log.Printf("Listening on %s", port)
	// register service

	// need to test find a way to implement ssl certificates
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	Bserver := server.BServer{}

	b_protopb.RegisterBMicroserviceServer(grpcServer, &Bserver)
	grpcServer.Serve(lis)
}
