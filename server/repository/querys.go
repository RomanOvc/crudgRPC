package repository

import (
	"context"
	pb "crudgRPC/proto"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Server struct {
	pb.UnimplementedUserCrudMnagmentServer
	Db *sqlx.DB
}

func NewgRPCServer(Db *sqlx.DB) *Server {
	return &Server{Db: Db}
}

func (s *Server) GetAll(ctx context.Context, in *pb.Empty) (*pb.CustomerList, error) {
	rows, err := s.Db.Query("SELECT customer_id, name, age, address FROM customer")
	if err != nil {
		return nil, errors.Wrap(err, "no data")
	}
	defer rows.Close()

	customers := []*pb.Customer{}
	for rows.Next() {
		var lc pb.Customer
		err := rows.Scan(&lc.CustomerId, &lc.Name, &lc.Age, &lc.Address)
		if err != nil {
			return nil, errors.Wrap(err, "no data")
		}
		customers = append(customers, &lc)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error row")
	}

	return &pb.CustomerList{Customer: customers}, nil
}

func (s *Server) GetByIdCustomer(ctx context.Context, in *pb.CustomerRequestId) (*pb.Customer, error) {
	var user pb.Customer
	rows := s.Db.QueryRow("select customer_id,name,age,address from customer where customer_id=$1", in.CustomerId)
	err := rows.Scan(&user.CustomerId, &user.Name, &user.Age, &user.Address)
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "customer_id not valid!")
	}
	if user.CustomerId == "" {
		return nil, errors.New("nno customer_id")
	}
	return &user, err
}

func (s *Server) InsertCustomer(ctx context.Context, in *pb.ICustomer) (*pb.StateMessage, error) {
	var id int64
	err := s.Db.QueryRow("insert into customer (name,age,address) values($1, $2, $3) RETURNING customer_id", in.Name, in.Age, in.Address).Scan(&id)
	if err != nil {
		return &pb.StateMessage{State: 0}, errors.Wrap(err, "customer not found")
	}
	if err != nil {
		return &pb.StateMessage{State: 0}, errors.Wrap(err, "customer not found")
	}
	if err != nil {
		return &pb.StateMessage{State: 0}, errors.Wrap(err, "customer is not add")
	}

	return &pb.StateMessage{State: id}, nil
}

func (s *Server) RemoveCustomer(ctx context.Context, in *pb.CustomerRequestId) (*pb.StateMessage, error) {
	var id int64
	err := s.Db.QueryRow("DELETE FROM customer where customer_id = $1 RETURNING customer_id ", in.CustomerId).Scan(&id)
	if err != nil {
		return &pb.StateMessage{State: 0}, errors.Wrap(err, "error username")
	}
	return &pb.StateMessage{State: id}, nil
}

func (s *Server) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	var customer pb.Customer
	log.Println(in.Name, in.Age, in.Address)
	err := s.Db.QueryRow("UPDATE customer SET name=$1, age=$2, address=$3 where customer_id=$4 RETURNING customer_id,name,age,address", in.Name, in.Age, in.Address, in.CustomerId).Scan(&customer.CustomerId, &customer.Name, &customer.Age, &customer.Address)
	if err != nil {
		return nil, errors.Wrap(err, "customer not found")
	}
	log.Println(&customer)
	return &customer, err

}
