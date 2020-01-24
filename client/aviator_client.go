package client

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"tix-proxy/domain"
)

type AviatorServiceClient interface {
	GetEventReviews()
	GetEventPhotos()
	GetEventAvailableDate()
	GetEventDetail()
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
	url, err := url.Parse(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	apiEndPoint := client.AviatorBaseURL.ResolveReference(url).String()

	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	query := req.URL.Query()
	query.Add("code", code)
	query.Add("sortOrder", sort)
	query.Add("topX", page)
	query.Add("showUnavailable", "false")
	query.Add("apiKey", client.AviatorApiKey)
	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}

	response, err := httpClient.Do(req)
	if err != nil {
		error := fmt.Errorf("The HTTP request failed with error %s", err.Error())
		log.Error(error)
		return nil, error
	}

	defer response.Body.Close()
	var reviews domain.AviatorReviewList

	if err := json.NewDecoder(response.Body).Decode(&reviews); err != nil {
		log.Error("[AVIATOR] Error when decode response")
		return nil, err
	}

	log.Print("[AVIATOR] Review Response Success")
	return &reviews, nil
}

func (client *AviatorClient) GetEventPhotos(code, sort string, start, finish int) (*domain.AviatorPhotoList, error) {
	page := fmt.Sprintf("%d-%d", start, finish)
	path := "/service/product/photos"
	url, err := url.Parse(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	apiEndPoint := client.AviatorBaseURL.ResolveReference(url).String()

	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	query := req.URL.Query()
	query.Add("code", code)
	query.Add("sortOrder", sort)
	query.Add("topX", page)
	query.Add("showUnavailable", "false")
	query.Add("apiKey", client.AviatorApiKey)
	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}

	response, err := httpClient.Do(req)
	if err != nil {
		error := fmt.Errorf("The HTTP request failed with error %s", err.Error())
		log.Error(error)
		return nil, error
	}

	defer response.Body.Close()
	var photos domain.AviatorPhotoList

	if err := json.NewDecoder(response.Body).Decode(&photos); err != nil {
		log.Error("[AVIATOR] Error when decode response")
		return nil, err
	}

	log.Info("[AVIATOR] Request Photo List Success")

	return &photos, nil
}

func (client *AviatorClient) GetEventAvailableDate(code string) (*domain.AviatorAvailableDate, error) {
	path := "/service/booking/availability/dates"
	url, err := url.Parse(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	apiEndPoint := client.AviatorBaseURL.ResolveReference(url).String()
	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	query := req.URL.Query()
	query.Add("productCode", code)
	query.Add("apiKey", client.AviatorApiKey)
	req.URL.RawQuery = query.Encode()

	LogRequest("[AVIATOR]", req)

	httpClient := &http.Client{}

	response, err := httpClient.Do(req)
	if err != nil {
		error := fmt.Errorf("The HTTP request failed with error %s", err.Error())
		log.Error(error)
		return nil, error
	}

	defer response.Body.Close()

	var date domain.AviatorAvailableDate
	if err := json.NewDecoder(response.Body).Decode(&date); err != nil {
		log.Error("[AVIATOR] Error when decode response")
		log.Error(err)
		return nil, err
	}

	log.Info("[AVIATOR] Request Available Date Success")

	return &date, nil
}

func (client *AviatorClient) GetEventDetail(code string) (*domain.AviatorProduct, error) {
	path := "/service/product"
	url, err := url.Parse(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	apiEndPoint := client.AviatorBaseURL.ResolveReference(url).String()
	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	query := req.URL.Query()
	query.Add("code", code)
	query.Add("apiKey", client.AviatorApiKey)
	req.URL.RawQuery = query.Encode()

	LogRequest("[AVIATOR]", req)

	httpClient := &http.Client{}

	response, err := httpClient.Do(req)
	if err != nil {
		error := fmt.Errorf("The HTTP request failed with error %s", err.Error())
		log.Error(error)
		return nil, error
	}

	defer response.Body.Close()

	var event domain.AviatorProduct
	if err := json.NewDecoder(response.Body).Decode(&event); err != nil {
		log.Error("[AVIATOR] Error when decode response")
		log.Error(err)
		return nil, err
	}

	log.Info("[AVIATOR] Request Event Detail Success")

	return &event, nil
}
