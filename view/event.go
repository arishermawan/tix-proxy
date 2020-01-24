package view

type EventResponse struct {
	Name              string           `json:"name"`
	City              string           `json:"city"`
	Overview          string           `json:"overview"`
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
