package caller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Caller struct {
	Client *http.Client
}

func (caller *Caller) Get(ctx context.Context, url string, queryMap map[string]string) (map[string]interface{}, error) {
	maxConnWaitTime := caller.Client.Timeout
	ctx, cancel := context.WithTimeout(ctx, maxConnWaitTime)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("error creating request:%v", err)
		return nil, err
	}

	// appending to existing query args
	query := req.URL.Query()
	for key, val := range queryMap {
		query.Add(key, val)
	}
	req.URL.RawQuery = query.Encode()

	res, err := caller.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("Error close body response: %v", err)
		}
	}()
	var apiResponse map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}

func (caller *Caller) SendRequest(ctx context.Context, method string, url string, queryMap map[string]string, headerMap map[string]string,
	bodyRequest interface{}) (map[string]interface{}, error) {

	maxConnWaitTime := caller.Client.Timeout
	ctx, cancel := context.WithTimeout(ctx, maxConnWaitTime)
	defer cancel()

	jsonData, err := json.Marshal(bodyRequest)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("error when create new request: %v", err)
		return nil, err
	}

	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	query := req.URL.Query()
	for key, val := range queryMap {
		query.Add(key, val)
	}
	req.URL.RawQuery = query.Encode()

	res, err := caller.Client.Do(req)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v", err)
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("Error close body response: %v", err)
		}
	}()
	var apiResponse map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}
