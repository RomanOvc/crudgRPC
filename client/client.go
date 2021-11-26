package main

import (
	"log"
	"net/http"

	"crudgRPC/client/handlers"
	pb "crudgRPC/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("did not connection :&w", err)
	}
	defer conn.Close()

	c := pb.NewUserCrudMnagmentClient(conn)
	handler := handlers.NewClientHandler(c)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get_all", handler.GetAllC).Methods("Get")
	router.HandleFunc("/get_by_id", handler.GetByIdC).Methods("Get")
	router.HandleFunc("/create_customer", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/delete_by_customer_id", handler.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/update_by_customer_id", handler.UpateCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))

}
