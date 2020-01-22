package view

import "tix-proxy/domain"

type AvailableDateResponse struct {
	AvailableDate map[string]interface{} `json:"available_dates"`
}

func NewAvailableDateResponse(date *domain.AviatorAvailableDate) AvailableDateResponse {
	return AvailableDateResponse{
		AvailableDate: date.AvailableDate,
	}
}
