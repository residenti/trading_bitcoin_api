package bitflyer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/residenti/trading_bitcoin_api/config"
)

type APIClient struct {
	httpClient *http.Client
}

func New() *APIClient {
	apiClient := &APIClient{&http.Client{}}
	return apiClient
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	httpBaseURL, err := url.Parse(config.List.HttpBaseUrl)
	if err != nil {
		log.Printf("Faild to parse httpBaseURL: %v", err)
	}

	path, err := url.Parse(urlPath)
	if err != nil {
		log.Printf("Faild to parse urlPath: %v", err)
	}

	endpoint := httpBaseURL.ResolveReference(path).String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Faild to http.NewRequest: %v", err)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (api *APIClient) GetTicker(productCode string) (*Ticker, error) {
	resp, err := api.doRequest("GET", "ticker", map[string]string{"product_code": productCode}, nil)
	if err != nil {
		return nil, err
	}

	var ticker Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, err
	}

	return &ticker, nil
}
