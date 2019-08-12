package user

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc/common/config"
	"log"
	"net/http"

	"github.com/muhammadhidayah/inact-grpc/client/middleware"
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
	"google.golang.org/grpc"
)

type UserHandler struct{}

func NewUserHandler(mux *middleware.GoMiddleware) {
	uh := &UserHandler{}
	customMiddleware := new(middleware.MiddlewareFunc)
	mux.HandleFunc("/auth/login", customMiddleware.CustomizeMiddleware(uh.DoLogin, customMiddleware.AllowMethod("POST")))

	mux.HandleFunc("/user", customMiddleware.CustomizeMiddleware(uh.GetUserID, customMiddleware.AllowMethod("GET")))

}

func (uh *UserHandler) GetUserID(w http.ResponseWriter, r *http.Request) {
	user := uh.ServiceUser()
	res, err := user.GetPersonByID(context.Background(), &_models.PersonID{
		PersonId: "hidayah@inarts.co.id",
	})

	fmt.Println(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	w.Write([]byte("\n"))

}

func (uh *UserHandler) ServiceUser() _models.PersonsClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to ", port, err)
	}

	return _models.NewPersonsClient(conn)
}

func (uh *UserHandler) DoLogin(w http.ResponseWriter, r *http.Request) {

}
