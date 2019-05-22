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
	r.Methods("OPTIONS").HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			headers := res.Header()
			headers.Add("Access-Control-Allow-Origin", "*")
			headers.Add("Vary", "Origin")
			headers.Add("Vary", "Access-Control-Request-Method")
			headers.Add("Vary", "Access-Control-Request-Headers")
			headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, az-auth-token")
			headers.Add("Access-Control-Allow-Methods", "GET,PUT,POST,OPTIONS")
		})

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
	r.HandleFunc("/users/{id}/assessments", wrap(s.getAssessments, s.authenticate, s.logIt, s.recover)).Methods("GET")
	r.HandleFunc("/users/{id}/assessments", wrap(s.postAssessment, s.authenticate, s.logIt, s.recover)).Methods("POST")
	r.HandleFunc("/users/{id}/assessments/{assessmentId}", wrap(s.putAssessment, s.authenticate, s.logIt, s.recover)).Methods("PUT")

	return s
}

func (s *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), s.router)
}
