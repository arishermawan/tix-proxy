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

type AviatorProduct struct {
	*AviatorGenericResponse
	Event AviatorEventDetail `json:"data"`
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

type AviatorEventDetail struct {
	SupplierName                  string                 `json:"supplierName"`
	CurrencyCode                  string                 `json:"currencyCode"`
	CatIds                        []int                  `json:"catIds"`
	SubCatIds                     []int                  `json:"subCatIds"`
	WebURL                        interface{}            `json:"webURL"`
	SpecialReservationDetails     interface{}            `json:"specialReservationDetails"`
	SslSupported                  bool                   `json:"sslSupported"`
	PanoramaCount                 int                    `json:"panoramaCount"`
	MerchantCancellable           bool                   `json:"merchantCancellable"`
	BookingEngineID               string                 `json:"bookingEngineId"`
	OnRequestPeriod               interface{}            `json:"onRequestPeriod"`
	PrimaryGroupID                interface{}            `json:"primaryGroupId"`
	Pas                           interface{}            `json:"pas"`
	Available                     bool                   `json:"available"`
	ProductURLName                string                 `json:"productUrlName"`
	PrimaryDestinationURLName     string                 `json:"primaryDestinationUrlName"`
	VoucherRequirements           string                 `json:"voucherRequirements"`
	TourGradesAvailable           bool                   `json:"tourGradesAvailable"`
	HotelPickup                   bool                   `json:"hotelPickup"`
	UserPhotos                    []AviatorPhoto         `json:"userPhotos"`
	Reviews                       []AviatorReview        `json:"reviews"`
	Videos                        interface{}            `json:"videos"`
	TourGrades                    []AviatorTourGrades    `json:"tourGrades"`
	AgeBands                      []AviatorAgeBands      `json:"ageBands"`
	BookingQuestions              []interface{}          `json:"bookingQuestions"`
	PassengerAttributes           interface{}            `json:"passengerAttributes"`
	Highlights                    interface{}            `json:"highlights"`
	SalesPoints                   []string               `json:"salesPoints"`
	RatingCounts                  map[string]int         `json:"ratingCounts"`
	TermsAndConditions            interface{}            `json:"termsAndConditions"`
	MerchantTermsAndConditions    AviatorMerchantTnc     `json:"merchantTermsAndConditions"`
	MaxTravellerCount             int                    `json:"maxTravellerCount"`
	Exclusions                    []string               `json:"exclusions"`
	ProductPhotos                 []AviatorProductPhotos `json:"productPhotos"`
	Itinerary                     string                 `json:"itinerary"`
	DestinationID                 int                    `json:"destinationId"`
	VoucherOption                 string                 `json:"voucherOption"`
	ApplePassSupported            bool                   `json:"applePassSupported"`
	TranslationLevel              int                    `json:"translationLevel"`
	AdditionalInfo                []string               `json:"additionalInfo"`
	Inclusions                    []string               `json:"inclusions"`
	DepartureTime                 string                 `json:"departureTime"`
	DepartureTimeComments         string                 `json:"departureTimeComments"`
	DeparturePoint                string                 `json:"departurePoint"`
	City                          string                 `json:"city"`
	ReturnDetails                 string                 `json:"returnDetails"`
	SpecialOffer                  string                 `json:"specialOffer"`
	Operates                      string                 `json:"operates"`
	MapURL                        interface{}            `json:"mapURL"`
	AllTravellerNamesRequired     bool                   `json:"allTravellerNamesRequired"`
	Description                   string                 `json:"description"`
	Location                      string                 `json:"location"`
	Country                       string                 `json:"country"`
	Region                        string                 `json:"region"`
	ShortDescription              string                 `json:"shortDescription"`
	Price                         float64                `json:"price"`
	SupplierCode                  string                 `json:"supplierCode"`
	ThumbnailHiResURL             string                 `json:"thumbnailHiResURL"`
	OnSale                        bool                   `json:"onSale"`
	PhotoCount                    int                    `json:"photoCount"`
	ReviewCount                   int                    `json:"reviewCount"`
	PrimaryDestinationID          int                    `json:"primaryDestinationId"`
	MerchantNetPriceFrom          float64                `json:"merchantNetPriceFrom"`
	PrimaryDestinationName        string                 `json:"primaryDestinationName"`
	ThumbnailURL                  string                 `json:"thumbnailURL"`
	PriceFormatted                string                 `json:"priceFormatted"`
	Rrp                           float64                `json:"rrp"`
	Rrpformatted                  string                 `json:"rrpformatted"`
	VideoCount                    int                    `json:"videoCount"`
	Rating                        float64                `json:"rating"`
	SpecialOfferAvailable         bool                   `json:"specialOfferAvailable"`
	SpecialReservation            bool                   `json:"specialReservation"`
	ShortTitle                    string                 `json:"shortTitle"`
	MerchantNetPriceFromFormatted string                 `json:"merchantNetPriceFromFormatted"`
	SavingAmount                  float64                `json:"savingAmount"`
	SavingAmountFormated          string                 `json:"savingAmountFormated"`
	Essential                     interface{}            `json:"essential"`
	Admission                     interface{}            `json:"admission"`
	Title                         string                 `json:"title"`
	Duration                      string                 `json:"duration"`
	Code                          string                 `json:"code"`
}

type AviatorTourGrades struct {
	SortOrder    int    `json:"sortOrder"`
	CurrencyCode string `json:"currencyCode"`
	LangServices struct {
		EnSERVICEAUDIO string `json:"en/SERVICE_AUDIO"`
	} `json:"langServices"`
	GradeCode                     string  `json:"gradeCode"`
	MerchantNetPriceFrom          float64 `json:"merchantNetPriceFrom"`
	PriceFrom                     float64 `json:"priceFrom"`
	PriceFromFormatted            string  `json:"priceFromFormatted"`
	MerchantNetPriceFromFormatted string  `json:"merchantNetPriceFromFormatted"`
	GradeTitle                    string  `json:"gradeTitle"`
	GradeDepartureTime            string  `json:"gradeDepartureTime"`
	GradeDescription              string  `json:"gradeDescription"`
	DefaultLanguageCode           string  `json:"defaultLanguageCode"`
}

type AviatorAgeBands struct {
	SortOrder         int    `json:"sortOrder"`
	AgeFrom           int    `json:"ageFrom"`
	AgeTo             int    `json:"ageTo"`
	BandID            int    `json:"bandId"`
	Adult             bool   `json:"adult"`
	PluralDescription string `json:"pluralDescription"`
	TreatAsAdult      bool   `json:"treatAsAdult"`
	Count             int    `json:"count"`
	Description       string `json:"description"`
}

type AviatorProductPhotos struct {
	PhotoURL string `json:"photoURL"`
	Supplier string `json:"supplier"`
	Caption  string `json:"caption"`
}

type AviatorCancellationFromTourDate struct {
	DayRangeMin          int         `json:"dayRangeMin"`
	DayRangeMax          interface{} `json:"dayRangeMax"`
	PercentageRefundable int         `json:"percentageRefundable"`
	PolicyStartTimestamp interface{} `json:"policyStartTimestamp"`
	PolicyEndTimestamp   interface{} `json:"policyEndTimestamp"`
}

type AviatorMerchantTnc struct {
	TermsAndConditions             string                            `json:"termsAndConditions"`
	MerchantTermsAndConditionsType int                               `json:"merchantTermsAndConditionsType"`
	AmountRefundable               interface{}                       `json:"amountRefundable"`
	CancellationFromTourDate       []AviatorCancellationFromTourDate `json:"cancellationFromTourDate"`
}
