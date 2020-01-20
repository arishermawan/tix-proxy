package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"tix-proxy/domain"
)

type AviatorServiceClient interface {
	GetEventReviews()
}

type AviatorClient struct {
	AviatorApiKey  string
	AviatorBaseURL *url.URL
}

func NewAviatorServiceClient(apiKey string, host *url.URL) *AviatorClient {
	return &AviatorClient{
		AviatorApiKey:  apiKey,
		AviatorBaseURL: host,
	}
}

func (client *AviatorClient) GetEventReviews(code, sort string, start, finish int) (*domain.AviatorReviewList, error) {
	page := fmt.Sprintf("%d-%d", start, finish)
	path := "/service/product/reviews"
	relativeURL := fmt.Sprintf("%s?code=%s&topX=%s&sortOrder=%s&showUnavailable=false&apiKey=%s", path, code, page, sort, client.AviatorApiKey)
	url, err := url.Parse(relativeURL)
	if err != nil {
		return nil, err
	}

	apiEndPoint := client.AviatorBaseURL.ResolveReference(url).String()
	response, err := http.Get(apiEndPoint)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s", err.Error())
	}

	// data, _ := ioutil.ReadAll(response.Body)
	var reviews domain.AviatorReviewList

	if err := json.NewDecoder(response.Body).Decode(&reviews); err != nil {
		return nil, err
	}

	// json.Unmarshal([]byte(data), &reviews)

	// userDetails := domain.UserDetails{}
	// if err := json.NewDecoder(resp.Body).Decode(&userDetails); err != nil {
	// 	log.WithField("error", err.Error()).Errorf("[UserServiceClient] parse user-details response")
	// 	return nil, err
	// }
	return &reviews, nil
}
