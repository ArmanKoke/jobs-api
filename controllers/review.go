package controllers

import (
	"context"
	"encoding/json"
	"jobs-api/models"
	"jobs-api/requests"
	"jobs-api/responses"
	"jobs-api/utils/databases"
	"jobs-api/utils/logger"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ReviewController struct
type ReviewController struct {
	log   logger.LogEnv
	dbEnv databases.DBEnv
}

// NewReviewController instance
func NewReviewController(log logger.LogEnv, dbEnv databases.DBEnv) *ReviewController {
	return &ReviewController{log, dbEnv}
}

// Create review
// swagger:route POST /review/create application CreateRequest
// Creates new review, see below.
// Responses:
//   200: body:CreateResponse
func (rc *ReviewController) Create(w http.ResponseWriter, r *http.Request) {
	zapLog := rc.log.ZapLogger()
	defer zapLog.Sync()

	var decodedReview requests.Review

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(&decodedReview)
	if err != nil {
		log.Println(err)
	}

	zapLog.Info("Request data", zap.String("ReviewID", decodedReview.ID))

	decodedReview.Validate()

	var created responses.Create
	created.Saved, err = models.Create(ctx, rc.log, rc.dbEnv, decodedReview)
	if err != nil {
		log.Println(err)
	}

	zapLog.Info("Reponse data", zap.Any("data", created))

	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		log.Printf("Error in review.Create method: %s", err)
	}
}

// Show reviews
// swagger:route GET /reviews application Show
// Shows reviews, see below.
// Responses:
//   200: body:ShowResponse
func (rc *ReviewController) Show(w http.ResponseWriter, r *http.Request) {
	zapLog := rc.log.ZapLogger()
	defer zapLog.Sync()

	reviews := make([]responses.Show, 0)
	var err error
	var review requests.Review

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	review.Value = r.URL.Query().Get("value")

	zapLog.Info("Request data", zap.String("value", review.Value))

	if review.Value != "" {
		reviews, err = models.Show(ctx, rc.log, rc.dbEnv, review.Value)
		if err != nil {
			log.Println(err)
		}
	}

	zapLog.Info("Response data", zap.Any("data", reviews))

	err = json.NewEncoder(w).Encode(reviews)
	if err != nil {
		log.Printf("Error in review.Show method: %s", err)
	}
}
