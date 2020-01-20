package view

import "tix-proxy/domain"

type ReviewResponse struct {
	Position      int    `json:"position"`
	Rating        int    `json:"rating"`
	Reviewer      string `json:"reviewer"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PublishedDate string `json:"published_date"`
}

func NewReviewResponse(review domain.AviatorReview) ReviewResponse {
	return ReviewResponse{
		Position:      review.SortOrder,
		Rating:        review.Rating,
		Reviewer:      review.OwnerName,
		Title:         review.ProductTitle,
		Description:   review.Review,
		PublishedDate: review.PublishedDate}
}
