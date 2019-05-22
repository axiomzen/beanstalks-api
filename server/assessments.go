package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/axiomzen/beanstalks-api/model"
	"github.com/gorilla/mux"
)

type getAssessmentsResponse struct {
	User        *model.User        `json:"user"`
	Assessments *model.Assessments `json:"assessments"`
}

func (s *Server) getAssessments(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		s.log.WithError(err).Error("invalid id in request")
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

type scorePayload struct {
	TrackID int `sql:"track_id,pk" json:"trackId"`
	StageID int `sql:"stage_id,pk" json:"stageId"`
	Score   int `json:"score"`
}

type assessmentPayload struct {
	ReviewerID int    `json:"reviewerId"`
	State      string `json:"state"`

	Scores []*scorePayload `json:"scores"`
}

type postAssessmentRequest struct {
	Assessment *assessmentPayload `json:"assessment"`
}

func (s *Server) postAssessment(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		s.log.WithError(err).Error("invalid id in request")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	assessmentReq := &postAssessmentRequest{}
	if err := json.NewDecoder(req.Body).Decode(assessmentReq); err != nil {
		s.log.WithError(err).Error("failed to decode assessment in request body")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// create the assessment
	assessment := &model.Assessment{
		UserID:     userID,
		ReviewerID: assessmentReq.Assessment.ReviewerID,
		State:      assessmentReq.Assessment.State,
	}
	if err := s.dal.CreateAssessment(assessment); err != nil {
		s.log.WithError(err).Error("failed to insert new assessment")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// get assessment so we have its ID
	if err := s.dal.GetAssessmentByPK(assessment); err != nil {
		s.log.WithError(err).Error("failed to fetch newly created assessment")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create scores for assessment
	for _, score := range assessmentReq.Assessment.Scores {
		if err := s.dal.CreateScore(&model.Score{
			AssessmentID: assessment.ID,
			TrackID:      score.TrackID,
			StageID:      score.StageID,
			Score:        score.Score,
		}); err != nil {
			s.log.WithError(err).Error("failed to create score for new assessment")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	res.WriteHeader(http.StatusCreated)
}
