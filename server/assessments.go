package server

import (
	"encoding/json"
	"net/http"

	"github.com/axiomzen/beanstalks-api/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type getAssessmentsResponse struct {
	User        *model.User        `json:"user"`
	Assessments *model.Assessments `json:"assessments"`
}

func (s *Server) getAssessments(res http.ResponseWriter, req *http.Request) {
	claims := req.Context().Value(claimsCtxKey).(jwt.MapClaims)
	user := &model.User{
		ID: int(claims["id"].(float64)),
	}

	// get user from DB
	if err := s.dal.GetUserByID(user); err != nil {
		s.log.WithError(err).Error("failed to fetch user from DB")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get all assessments for the user
	assessments := &model.Assessments{}
	if err := s.dal.GetAssessmentsByUserID(user.ID, assessments); err != nil {
		s.log.WithError(err).Error("failed to fetch assessments for user")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := &getAssessmentsResponse{
		User:        user,
		Assessments: assessments,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(payload)
}
