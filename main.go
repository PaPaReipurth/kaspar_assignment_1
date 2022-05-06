package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/papareipurth/kaspar_assignment_1/link_station"
	"github.com/papareipurth/kaspar_assignment_1/server"
	"google.golang.org/grpc"
)

var (
	port = *flag.Int("port", 9111, "Port on which to run the application")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Unable to serve with error: %v", err)
	}

	grpcServer := grpc.NewServer()
	link_station.RegisterLinkStationServiceServer(grpcServer, &server.LinkStationServer{})

	log.Printf("Listning on port %d...", port)
	grpcServer.Serve(lis)

}
