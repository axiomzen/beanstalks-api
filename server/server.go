package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/axiomzen/beanstalks-api/config"
	"github.com/axiomzen/beanstalks-api/data"
	"github.com/axiomzen/beanstalks-api/model"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Server struct {
	config *config.Config
	dal    *data.DAL
	router *mux.Router
	log    *logrus.Entry
}

func New(c *config.Config) *Server {
	r := mux.NewRouter()
	s := &Server{
		config: c,
		dal:    data.New(c),
		router: r,
		log:    logrus.NewEntry(logrus.StandardLogger()),
	}

	// Attatch request handlers
	r.HandleFunc("/api/users", s.getUser).Methods("GET")
	r.HandleFunc("/signin", s.Signin).Methods("POST")
	return s
}

func (s *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), s.router)
}

func (s *Server) getUser(res http.ResponseWriter, req *http.Request) {
	// TODO
	s.log.Infof("received request %v", req)
	res.WriteHeader(http.StatusOK)
}

func (s *Server) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := model.User{
		Email:          creds.Username,
		HashedPassword: creds.Password,
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(creds.Password))
	err = s.dal.GetUserByEmail(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
