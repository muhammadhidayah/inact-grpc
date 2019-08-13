package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"grpc/common/config"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/muhammadhidayah/inact-grpc/client/middleware"
	_models "github.com/muhammadhidayah/inact-grpc/common/models"
	"google.golang.org/grpc"
)

type UserHandler struct{}
type M map[string]interface{}

func NewUserHandler(mux *middleware.GoMiddleware) {
	uh := &UserHandler{}
	customMiddleware := new(middleware.MiddlewareFunc)

	mux.HandleFunc("/auth/login", customMiddleware.CustomizeMiddleware(uh.DoLogin, customMiddleware.AllowMethod("POST")))

	mux.HandleFunc("/user", customMiddleware.CustomizeMiddleware(uh.GetUserID, customMiddleware.AllowMethod("GET")))

	mux.HandleFunc("/users", customMiddleware.CustomizeMiddleware(uh.FetchAllUser, customMiddleware.AllowMethod("GET")))

	mux.HandleFunc("/postuser", customMiddleware.CustomizeMiddleware(uh.StoreUser, customMiddleware.AllowMethod("POST")))

	mux.HandleFunc("/updateuser", customMiddleware.CustomizeMiddleware(uh.StoreUser, customMiddleware.AllowMethod("POST")))

	mux.HandleFunc("/deleteuser", customMiddleware.CustomizeMiddleware(uh.DeleteUser, customMiddleware.AllowMethod("GET")))

}

func (uh *UserHandler) GetUserID(w http.ResponseWriter, r *http.Request) {
	user := uh.ServiceUser()
	res, err := user.GetPersonByID(context.Background(), &_models.PersonID{
		PersonId: r.FormValue("id"),
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

func (uh *UserHandler) FetchAllUser(w http.ResponseWriter, r *http.Request) {
	user := uh.ServiceUser()
	res, err := user.ListPerson(context.Background(), new(empty.Empty))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(res.List) > 0 {
		dataJson, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(dataJson)
	} else {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}
}

func (uh *UserHandler) StoreUser(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	data := _models.Person{}

	if err := payload.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	user := uh.ServiceUser()
	var err error
	if data.PersonNo != 0 {
		_, err = user.UpdatePerson(context.Background(), &data)
	} else {
		_, err = user.AddPerson(context.Background(), &data)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("success"))

}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	person_id := r.FormValue("id")
	data := _models.PersonID{
		PersonId: person_id,
	}

	user := uh.ServiceUser()

	_, err := user.DeletePerson(context.Background(), &data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("success"))
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
	username, password, isOk := r.BasicAuth()

	if !isOk {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err, userInfo := uh.authenticationUser(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// This is for token jwt set in localstorage
	signedToken, err := generateToken(false, userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, _ := json.Marshal(M{"token": signedToken})

	// This is for token jwt set in cookie
	signedTokenCookie, err := generateToken(true, userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieName := "RefreshTokenJWT"
	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = signedTokenCookie
		c.Expires = time.Now().Add(2 * time.Hour)
		c.HttpOnly = true
		http.SetCookie(w, c)
	}

	w.Write([]byte(tokenString))

}

func (uh *UserHandler) authenticationUser(username string, password string) (error, M) {
	user := uh.ServiceUser()
	person, err := user.GetPersonByID(context.Background(), &_models.PersonID{
		PersonId: username,
	})

	if err != nil {
		return err, nil
	}

	if person.PersonPassword != password {
		return errors.New("Your username or password is not valid"), nil
	}

	personNo := strconv.Itoa(int(person.PersonNo))

	dataM := M{
		"PersonNo": personNo,
		"PersonID": person.PersonId,
	}

	return nil, dataM
}

func generateToken(isLong bool, data M) (string, error) {
	duration := middleware.LOGIN_EXPIRATION_DURATION

	if isLong {
		duration = middleware.COOKIE_EXPIRATION_DURATION
	}

	claims := middleware.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    middleware.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		PersonNo: data["PersonNo"].(string),
		PersonID: data["PersonID"].(string),
	}

	token := jwt.NewWithClaims(
		middleware.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(middleware.JWT_SIGNATURE_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
