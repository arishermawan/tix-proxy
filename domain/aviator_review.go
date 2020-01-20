package domain

type AviatorReviewList struct {
	Success          bool            `json:"success"`
	Reviews          []AviatorReview `json:"data"`
	TotalCount       int             `json:"totalCount"`
	DateStamp        string          `json:"dateStamp"`
	ErrorReference   string          `json:"errorReference"`
	ErrorName        string          `json:"errorName"`
	ErrorType        string          `json:"errorType"`
	ErrorCodes       []string        `json:"errorCodes"`
	ErrorMessage     []string        `json:"errorMessage"`
	ErrorMessageText []string        `json:"errorMessageText"`
}

type AviatorReview struct {
	SortOrder      int    `json:"sortOrder"`
	OwnerID        int    `json:"ownerId"`
	ReviewID       int    `json:"reviewId"`
	Rating         int    `json:"rating"`
	OwnerName      string `json:"ownerName"`
	Review         string `json:"review"`
	OwnerCountry   string `json:"ownerCountry"`
	ProductTitle   string `json:"productTitle"`
	ProductURLName string `json:"productUrlName"`
	OwnerAvatarURL string `json:"ownerAvatarURL"`
	ProductCode    string `json:"productCode"`
	SubmissionDate string `json:"submissionDate"`
	PublishedDate  string `json:"publishedDate"`
	ViatorNotes    string `json:"viatorNotes"`
	ViatorFeedback string `json:"viatorFeedback"`
	SslSupported   bool   `json:"sslSupported"`
}
