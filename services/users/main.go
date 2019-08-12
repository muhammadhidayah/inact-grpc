package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/muhammadhidayah/inact-grpc/common/config"
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
	"github.com/muhammadhidayah/inact-grpc/services/users/user"
	"github.com/muhammadhidayah/inact-grpc/services/users/user/repository"
	"github.com/muhammadhidayah/inact-grpc/services/users/user/usecase"
	"google.golang.org/grpc"
)

func Connect() (*sql.DB, error) {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "docker"
	dbName := "inact-grpc"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

type UserServer struct{}

var userRepo user.Repository
var userUseCase user.Usecase

func init() {
	conn, err := Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo = repository.NewUserRepository(conn)
	userUseCase = usecase.NewUserUseCase(userRepo)
}

func (UserServer) GetPersonByID(ctx context.Context, param *_models.PersonID) (*_models.Person, error) {
	personId := param.PersonId
	res, err := userUseCase.GetUserByID(personId)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	srv := grpc.NewServer()
	var userSrv UserServer
	_models.RegisterPersonsServer(srv, userSrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)

	if err != nil {
		log.Fatalf("couldnt listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
