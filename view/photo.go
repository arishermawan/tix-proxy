package view

import "tix-proxy/domain"

type PhotoResponse struct {
	Position     int    `json:"position"`
	Thumbnail    string `json:"thumbnail"`
	Photo        string `json:"photo"`
	Title        string `json:"title"`
	TimeUploaded string `json:"time_uploaded"`
}

func NewPhotoResponse(photo domain.AviatorPhoto) PhotoResponse {
	return PhotoResponse{
		Position:     photo.SortOrder,
		Thumbnail:    photo.ThumbnailURL,
		Photo:        photo.PhotoURL,
		Title:        photo.Title,
		TimeUploaded: photo.TimeUploaded}
}
