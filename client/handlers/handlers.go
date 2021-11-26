package handlers

import (
	"context"
	pb "crudgRPC/proto"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ClientHandler struct {
	PB pb.UserCrudMnagmentClient
}

func NewClientHandler(PB pb.UserCrudMnagmentClient) *ClientHandler {
	return &ClientHandler{PB: PB}
}

func (chr *ClientHandler) GetAllC(w http.ResponseWriter, r *http.Request) {
	var (
		u   []byte
		err error
	)

	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.WriteHeader(400)
			w.Write(nil)
		} else {
			w.Write(u)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := chr.PB.GetAll(ctx, &pb.Empty{})

	u, err = json.Marshal(res)
	if err != nil {
		return
	}
	return
}

func (chr *ClientHandler) GetByIdC(w http.ResponseWriter, r *http.Request) {
	var (
		u   []byte
		err error
	)
	w.Header().Set("Content-Type", "applicaion/json")
	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.WriteHeader(400)
			w.Write(nil)
		} else {
			w.Write(u)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id_customer := r.URL.Query().Get("id_customer")
	res, err := chr.PB.GetByIdCustomer(ctx, &pb.CustomerRequestId{CustomerId: id_customer})
	if err != nil {
		return
	}
	u, err = json.Marshal(res)
	if err != nil {
		return
	}
	return
}

func (chr *ClientHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		bytes []byte
	)
	var customer *pb.Customer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/octet-stream")
	json.NewDecoder(r.Body).Decode(&customer)

	defer func() {
		if err != nil {
			w.WriteHeader(400)
			bytes, _ := json.Marshal("not" + err.Error())
			w.Write(bytes)
			return
		} else {
			w.Write(bytes)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customId, err := chr.PB.InsertCustomer(ctx, &pb.ICustomer{Name: customer.Name, Age: customer.Age, Address: customer.Address})
	log.Println(customId)
	bytes, err = json.Marshal(customId)
	if err != nil {
		return
	}
	return
}

func (chr *ClientHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		bytes []byte
	)
	w.Header().Set("Content-Type", "applicaion/json")

	defer func() {
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			w.Write(nil)
		} else {
			w.Write(bytes)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id_customer := r.URL.Query().Get("id_customer")
	log.Println(id_customer)
	customId, err := chr.PB.RemoveCustomer(ctx, &pb.CustomerRequestId{CustomerId: id_customer})
	bytes, err = json.Marshal(customId)
	if err != nil {
		return
	}
	return
}

func (chr *ClientHandler) UpateCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		bytes []byte
	)

	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			w.Write(nil)
		} else {
			w.Write(bytes)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var customer *pb.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	// log.Println(customer)
	newCustomer, err := chr.PB.UpdateCustomer(ctx, &pb.Customer{CustomerId: customer.CustomerId, Name: customer.Name, Age: customer.Age, Address: customer.Address})
	if err != nil {
		return
	}

	bytes, err = json.Marshal(newCustomer)
	if err != nil {
		return
	}
	return
}
