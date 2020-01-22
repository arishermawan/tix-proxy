package domain

type AviatorGenericResponse struct {
	Success          bool     `json:"success"`
	TotalCount       int      `json:"totalCount"`
	DateStamp        string   `json:"dateStamp"`
	ErrorReference   string   `json:"errorReference"`
	ErrorName        string   `json:"errorName"`
	ErrorType        string   `json:"errorType"`
	ErrorCodes       []string `json:"errorCodes"`
	ErrorMessage     []string `json:"errorMessage"`
	ErrorMessageText []string `json:"errorMessageText"`
}

type AviatorReviewList struct {
	*AviatorGenericResponse
	Reviews []AviatorReview `json:"data"`
}

type AviatorPhotoList struct {
	*AviatorGenericResponse
	Photos []AviatorPhoto `json:"data"`
}

type AviatorAvailableDate struct {
	*AviatorGenericResponse
	AvailableDate map[string]interface{} `json:"data"`
}

type Infos map[string][]string

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

type AviatorPhoto struct {
	SortOrder         int         `json:"sortOrder"`
	OwnerName         string      `json:"ownerName"`
	OwnerID           int         `json:"ownerId"`
	OwnerCountry      interface{} `json:"ownerCountry"`
	ProductTitle      string      `json:"productTitle"`
	ProductURLName    string      `json:"productUrlName"`
	OwnerAvatarURL    interface{} `json:"ownerAvatarURL"`
	SslSupported      bool        `json:"sslSupported"`
	ProductCode       string      `json:"productCode"`
	ThumbnailURL      string      `json:"thumbnailURL"`
	Caption           string      `json:"caption"`
	TimeUploaded      string      `json:"timeUploaded"`
	EditorsPick       bool        `json:"editorsPick"`
	PhotoID           int         `json:"photoId"`
	PhotoURL          string      `json:"photoURL"`
	PhotoHiResURL     string      `json:"photoHiResURL"`
	PhotoMediumResURL string      `json:"photoMediumResURL"`
	Title             string      `json:"title"`
}
