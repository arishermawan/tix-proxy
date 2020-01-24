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

func AviatorPhotoHandler(w http.ResponseWriter, r *http.Request) {
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

	photos, err := aviator.GetEventPhotos(eventCode, defaultRating, start, finish)
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusInternalServerError, err)
		return
	}

	if photos.Success {
		var photoListResponse []view.PhotoResponse
		for _, v := range photos.Photos {
			photoListResponse = append(photoListResponse, view.NewPhotoResponse(v))
		}
		ResponseRenderer(w, http.StatusOK, photoListResponse, true, nil)
	} else {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New(photos.ErrorMessageText[0]))
	}
}

func AviatorAvailableDateHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	eventCode := query.Get("external_event_code")

	aviator := client.NewAviatorServiceClient(config.AviatorApiKey(), config.AviatorBaseURL())

	dates, err := aviator.GetEventAvailableDate(eventCode)
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusInternalServerError, err)
		return
	}

	if dates.Success {
		availableDateResponse := view.NewAvailableDateResponse(dates)
		ResponseRenderer(w, http.StatusOK, availableDateResponse, true, nil)
	} else {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New(dates.ErrorMessageText[0]))
	}
}

func AviatorEventDetailHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	eventCode := query.Get("external_event_code")

	aviator := client.NewAviatorServiceClient(config.AviatorApiKey(), config.AviatorBaseURL())

	eventDetail, err := aviator.GetEventDetail(eventCode)
	if err != nil {
		RawErrorResponseRenderer(w, http.StatusInternalServerError, err)
		return
	}

	if eventDetail.Success {
		// availableDateResponse := view.NewAvailableDateResponse(dates)
		ResponseRenderer(w, http.StatusOK, eventDetail.Event, true, nil)
	} else {
		RawErrorResponseRenderer(w, http.StatusUnprocessableEntity, errors.New(eventDetail.ErrorMessageText[0]))
	}
}

func convertPageToRange(pageNo, pageSize int) (start, finish int) {
	start = pageNo*pageSize - (pageSize - 1)
	finish = pageNo * pageSize
	return
}
