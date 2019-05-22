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
	r.HandleFunc("/me", wrap(s.me, s.authenticate, s.logIt, s.recover)).Methods("GET")
	r.HandleFunc("/signup", wrap(s.signUp, s.logIt, s.recover)).Methods("POST")
	r.HandleFunc("/signin", wrap(s.signIn, s.logIt, s.recover)).Methods("POST")
	r.HandleFunc("/assessments", wrap(s.getAssessments, s.authenticate, s.logIt, s.recover)).Methods("GET")
	r.HandleFunc("/assessments", wrap(s.postAssessment, s.authenticate, s.logIt, s.recover)).Methods("POST")

	return s
}

func (s *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), s.router)
}
