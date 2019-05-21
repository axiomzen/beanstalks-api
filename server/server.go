package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/axiomzen/beanstalks-api/config"
	"github.com/axiomzen/beanstalks-api/data"
	"github.com/gorilla/mux"
)

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
	r.HandleFunc("/signin", data.Signin).Methods("POST")
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
