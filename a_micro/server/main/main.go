package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc_micro/a_micro/a_protopb"
	"grpc_micro/a_micro/server"
	pb "grpc_micro/b_micro/client/grpc"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "8081", "port number")
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

	// get the dependency client
	bc, err := pb.NewBClient(":8080")
	if err != nil {
		log.Println(err)
	}

	Aserver := server.AServer{Bc: bc}

	a_protopb.RegisterAMicroserviceServer(grpcServer, &Aserver)
	grpcServer.Serve(lis)
}
