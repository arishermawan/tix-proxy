package view

import "tix-proxy/domain"

type EventResponse struct {
	Name              string           `json:"name"`
	City              string           `json:"city"`
	SalesPoints       []string         `json:"sales_points"`
	AdditionalInfo    []string         `json:"additional_info"`
	Reviews           []ReviewResponse `json:"reviews"`
	Photos            []PhotoResponse  `json:"photos"`
	ProductPhotos     []ProductPhotos  `json:"product_photos"`
	TermsAndCondition interface{}      `json:"terms_and_condition"`
	Inclusions        []string         `json:"inclusions"`
	Exclusions        []string         `json:"exclusions"`
	DepartureTime     string           `json:"departure_time"`
	DeparturePoint    string           `json:"departure_point"`
}

type ProductPhotos struct {
	Photo    string `json:"photo"`
	Caption  string `json:"caption"`
	Supplier string `json:"supplier"`
}

func NewEventResponse(product domain.AviatorEventDetail) EventResponse {
	return EventResponse{
		Name:              product.SupplierName,
		City:              product.City,
		SalesPoints:       product.SalesPoints,
		AdditionalInfo:    product.AdditionalInfo,
		Reviews:           NewReviewListResponse(product.Reviews),
		Photos:            NewPhotoListResponse(product.UserPhotos),
		TermsAndCondition: product.TermsAndConditions,
		Inclusions:        product.Inclusions,
		Exclusions:        product.Exclusions,
		DepartureTime:     product.DepartureTime,
		DeparturePoint:    product.DeparturePoint,
	}
}
