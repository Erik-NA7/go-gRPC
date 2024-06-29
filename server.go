package main

import (
	"fmt"
	"log"
	"net"
	"os"
	controllers "unsia/controller"
	"unsia/pb/cities"
	"unsia/pkg/database"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// ========================================================================
	// Logging

	log := log.New(os.Stdout, "GoGRPC : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := database.OpenDB()
	if err != nil {
			log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	cityServer := controllers.City{DB: db, Log: log}

	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)

	fmt.Println("running grpc server...")

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("fialed to serve: %s", err)
		return
	}
}
