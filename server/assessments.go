package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/axiomzen/beanstalks-api/model"
)

type getAssessmentsResponse struct {
	User        *model.User        `json:"user"`
	Assessments *model.Assessments `json:"assessments"`
}

func (s *Server) getAssessments(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	if err != nil {
		s.log.WithError(err).Error("invalid user_id in request")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// get user from DB
	user := &model.User{ID: userID}
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

	// get scores for each assessment
	for _, a := range *assessments {
		scores := &model.Scores{}
		if err := s.dal.GetScoresByAssessmentID(a.ID, scores); err != nil {
			s.log.WithError(err).Error("failed to get scores for assessments")
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		a.Scores = append(a.Scores, *scores...)
	}

	payload := &getAssessmentsResponse{
		User:        user,
		Assessments: assessments,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(payload)
}
