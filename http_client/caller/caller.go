package caller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Caller struct {
	Client *http.Client
}

func (caller *Caller) Get(ctx context.Context, url string) (map[string]interface{}, error) {
	maxConnWaitTime := caller.Client.Timeout
	ctx, cancel := context.WithTimeout(ctx, maxConnWaitTime)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("error creating request:%v", err)
		return nil, err
	}
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
