package main

import (
	"crudgRPC/server/postgres"
	"crudgRPC/server/repository"
	"log"
	"net"

	pb "crudgRPC/proto"

	"google.golang.org/grpc"
)

const (
	host     = "localhost"
	portpsql = "5433"
	username = "postgres"
	dbname   = "crudgrpc"
	sslmode  = "disable"
	password = "acer5800"
	portgRPC = ":50051"
)

func main() {
	db, err := postgres.InitPostgresDB(postgres.Config{
		Host:     host,
		Port:     portpsql,
		Username: username,
		DBName:   dbname,
		SSLMode:  sslmode,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", portgRPC)
	if err != nil {
		log.Fatal("filed to listen: %w", err)
	}

	s := grpc.NewServer()
	postgresgRPC := repository.NewgRPCServer(db)
	pb.RegisterUserCrudMnagmentServer(s, postgresgRPC)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
