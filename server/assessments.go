package server

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/axiomzen/beanstalks-api/model"
	"github.com/gorilla/mux"
)

type stageResPayload struct {
	ID          int    `sql:",pk" json:"id"`
	TrackID     int    `sql:"track_id" json:"trackId"`
	Description string `sql:"description" json:"description"`
	Level       int    `json:"level"`

	Score int `json:"score"`
}

type trackResPayload struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`

	Stages []*stageResPayload `json:"stages"`
}

type assessmentResPayload struct {
	ID        int         `json:"id"`
	UserID    int         `json:"userId"`
	Reviewer  *model.User `json:"reviewer"`
	State     string      `json:"state"`
	CreatedAt time.Time   `json:"createdAt"`

	Tracks []*trackResPayload `json:"tracks"`
}

type getAssessmentsResponse struct {
	User        *model.User            `json:"user"`
	Assessments []assessmentResPayload `json:"assessments"`
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

	// create response payload
	assessmentsPayloads := []assessmentResPayload{}
	for _, a := range *assessments {
		arp := assessmentResPayload{
			ID:        a.ID,
			UserID:    a.UserID,
			Reviewer:  a.Reviewer,
			State:     a.State,
			CreatedAt: a.CreatedAt,
		}

		// fill in tracks for assessment
		tracksByID := map[int]*trackResPayload{}

		for _, score := range a.Scores {
			// add all tracks
			track := &trackResPayload{
				ID:          score.Track.ID,
				Name:        score.Track.Name,
				Description: score.Track.Description,
				Tags:        score.Track.Tags,
			}
			tracksByID[track.ID] = track
		}

		for _, score := range a.Scores {
			// add stages for each track, with scores
			stage := &stageResPayload{
				ID:          score.Stage.ID,
				TrackID:     score.Stage.TrackID,
				Description: score.Stage.Description,
				Level:       score.Stage.Level,
				Score:       score.Score,
			}
			track := tracksByID[stage.TrackID]
			track.Stages = append(track.Stages, stage)
		}

		for _, track := range tracksByID {
			// sort the stages in the track add it to the assessment
			sort.Slice(track.Stages, func(i, j int) bool {
				return track.Stages[i].Level < track.Stages[j].Level
			})
			arp.Tracks = append(arp.Tracks, track)
		}

		assessmentsPayloads = append(assessmentsPayloads, arp)
	}

	payload := &getAssessmentsResponse{
		User:        user,
		Assessments: assessmentsPayloads,
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

	Scores []*scorePayload `json:"tracks"`
}

type postAssessmentRequest struct {
	Assessment *assessmentPayload `json:"assessment"`
}

func (s *Server) postAssessment(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		s.log.WithError(err).Error("invalid user ID in request")
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

func (s *Server) putAssessment(res http.ResponseWriter, req *http.Request) {
	assessmentID, err := strconv.Atoi(mux.Vars(req)["assessmentId"])
	if err != nil {
		s.log.WithError(err).Error("invalid assessment ID in request")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	assessmentReq := &postAssessmentRequest{}
	if err := json.NewDecoder(req.Body).Decode(assessmentReq); err != nil {
		s.log.WithError(err).Error("failed to decode assessment in request body")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// update the assessment
	assessment := &model.Assessment{
		ID:         assessmentID,
		ReviewerID: assessmentReq.Assessment.ReviewerID,
		State:      assessmentReq.Assessment.State,
	}

	if err := s.dal.UpdateAssessmentByPK(assessment); err != nil {
		s.log.WithError(err).Error("failed to create or update assessment")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update scores
	for _, score := range assessmentReq.Assessment.Scores {
		updatedScore := &model.Score{
			AssessmentID: assessmentID,
			TrackID:      score.TrackID,
			StageID:      score.StageID,
			Score:        score.Score,
		}
		if err := s.dal.UpdateScoreByPK(updatedScore); err != nil {
			s.log.WithError(err).Error("failed to update score for assessment")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	res.WriteHeader(http.StatusOK)
}
