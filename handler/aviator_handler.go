package handler

import (
	"errors"
	"net/http"
	"strconv"
	"tix-proxy/client"
	"tix-proxy/config"
	"tix-proxy/view"
)

func AviatorReviewHandler(w http.ResponseWriter, r *http.Request) {
	const defaultRating = "REVIEW_RATING_D"
	query := r.URL.Query()
	pageNo, err := strconv.Atoi(query.Get("page_no"))
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New("page_no be blank"))
		return
	}
	pageSize, err := strconv.Atoi(query.Get("page_size"))
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New("page_size cant be blank"))
		return
	}
	eventCode := query.Get("external_event_code")
	sortOrder := query.Get("sort_order")
	if sortOrder == "" {
		sortOrder = defaultRating
	}

	start, finish := convertPageToRange(pageNo, pageSize)

	aviator := client.NewAviatorServiceClient(config.AviatorApiKey(), config.AviatorBaseURL())

	reviews, err := aviator.GetEventReviews(eventCode, defaultRating, start, finish)
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusInternalServerError, err)
		return
	}

	if reviews.Success {
		var reviewListResponse []view.ReviewResponse
		for _, v := range reviews.Reviews {
			reviewListResponse = append(reviewListResponse, view.NewReviewResponse(v))
		}
		ResponseRenderer(w, http.StatusOK, reviewListResponse, true, nil)
	} else {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New(reviews.ErrorMessageText[0]))
	}
}

func convertPageToRange(pageNo, pageSize int) (start, finish int) {
	start = pageNo*pageSize - (pageSize - 1)
	finish = pageNo * pageSize
	return
}
