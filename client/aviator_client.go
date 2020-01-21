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
	GetEventPhotos()
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

	defer response.Body.Close()
	var reviews domain.AviatorReviewList

	if err := json.NewDecoder(response.Body).Decode(&reviews); err != nil {
		return nil, err
	}

	return &reviews, nil
}

func (client *AviatorClient) GetEventPhotos(code, sort string, start, finish int) (*domain.AviatorPhotoList, error) {
	page := fmt.Sprintf("%d-%d", start, finish)
	path := "/service/product/photos"
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

	defer response.Body.Close()
	var photos domain.AviatorPhotoList

	if err := json.NewDecoder(response.Body).Decode(&photos); err != nil {
		return nil, err
	}

	return &photos, nil
}
