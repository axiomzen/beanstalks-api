package server

import (
	"encoding/json"
	"net/http"

	"github.com/axiomzen/beanstalks-api/model"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type signUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	Token string `json:"token"`
}

func (s *Server) signUp(res http.ResponseWriter, req *http.Request) {
	payload := &signUpRequest{}
	json.NewDecoder(req.Body).Decode(payload)

	hashedPW, err := hashPassword(payload.Password)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password for new user")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Name:           payload.Name,
		Email:          payload.Email,
		HashedPassword: string(hashedPW),
	}
	if err := s.dal.CreateUser(user); err != nil {
		s.log.WithError(err).Error("failed to create new user")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the user so we have the ID we need to issue a token
	s.dal.GetUserByEmail(user)

	token, err := newSignedJWT(user.ID, "breanstalk", s.config.Secret)
	if err != nil {
		s.log.WithError(err).Error("Failed to generate JWT for new user")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Put the token in the response
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(&signUpResponse{Token: token})
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) signIn(res http.ResponseWriter, req *http.Request) {
	signInRequest := &signInRequest{}
	if err := json.NewDecoder(req.Body).Decode(signInRequest); err != nil {
		s.log.WithError(err).Error("failed to parse request body")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get user from db
	user := &model.User{Email: signInRequest.Email}
	if err := s.dal.GetUserByEmail(user); err != nil {
		s.log.WithError(err).Error("failed to fetch user from db")
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(signInRequest.Password)); err != nil {
		s.log.WithError(err).Error("error validating user password")
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	// return token
	token, err := newSignedJWT(user.ID, "breanstalk", s.config.Secret)
	if err != nil {
		s.log.WithError(err).Error("Failed to generate JWT for new user")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Put the token in the response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&signUpResponse{Token: token})
}

func (s *Server) me(res http.ResponseWriter, req *http.Request) {
	claims := req.Context().Value(claimsCtxKey).(jwt.MapClaims)
	user := &model.User{
		ID: int(claims["id"].(float64)),
	}

	if err := s.dal.GetUserByID(user); err != nil {
		s.log.WithError(err).Error("failed to fetch user from DB")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(user)
}
