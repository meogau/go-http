package main

import (
	"context"
	"fmt"
	"go-http/http_client/caller"
	"go-http/http_client/high_performance_client"
	"time"
)

func main() {
	client := high_performance_client.GetHttpClient()
	httpCaller := caller.Caller{
		Client: client,
	}
	for i := 0; i < 10; i++ {
		start := time.Now()
		mapResponse, _ := httpCaller.Get(context.Background(), "https://httpbin.org/get")
		for key, val := range mapResponse {
			fmt.Printf(" %v : %v\n", key, val)
		}
		fmt.Printf("Request %v - take time: %v ms\n", i, time.Since(start))
	}
}
