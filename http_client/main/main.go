package main

import (
	"context"
	"fmt"
	"go-http/http_client/caller"
	"go-http/http_client/high_performance_client"
	"net/http"
	"time"
)

func main() {
	client := high_performance_client.GetHttpClient()
	httpCaller := caller.Caller{
		Client: client,
	}

	//slow in first request
	start := time.Now()
	mapResponse, _ := httpCaller.Get(context.Background(), "https://httpbin.org/get", nil)
	fmt.Printf("First time request - takes time: %v ms\n", time.Since(start))
	for key, val := range mapResponse {
		fmt.Printf(" %v : %v\n", key, val)
	}
	//faster in the latter request cuz reuse connection
	for i := 0; i < 10; i++ {
		start := time.Now()
		mapResponse, _ = httpCaller.Get(context.Background(), "https://httpbin.org/get", nil)
		fmt.Printf("Request %v - takes time: %v ms\n", i+2, time.Since(start))
	}
	//test method post
	mapResponse, _ = httpCaller.SendRequest(context.Background(), http.MethodPost, "https://httpbin.org/post", nil, nil, map[string]string{"foo": "baz"})
	for key, val := range mapResponse {
		fmt.Printf(" %v : %v\n", key, val)
	}
}
