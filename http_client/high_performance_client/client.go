package high_performance_client

import (
	"net/http"
	"time"
)

// https://www.loginradius.com/blog/engineering/tune-the-go-http-client-for-high-performance/
func GetHttpClient() *http.Client {
	//By default, the Golang Http client performs the connection pooling.
	//When the request completes, that connection remains open until the idle connection timeout (default is 90 seconds).
	//If another request came, that uses the same established connection instead of creating a new connection, after the idle connection time, the connection will return to the pool.
	//Using the connection pooling will keep less connection open and more requests will be served with minimal server resources.
	//When we not defined transport in the http.Client, it uses the default transport Go HTTP Transport
	//Default configuration of the HTTP Transport,
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	//The HTTP client does not contain the request timeout setting by default.
	//If you are using http.Get(URL) or &Client{} that uses the http.DefaultClient.
	//DefaultClient has not timeout setting; it comes with no timeout
	//For the Rest API, it is recommended that timeout should not more than 10 seconds.
	//If the Requested resource is not responded to in 10 seconds,
	//the HTTP connection will be canceled with net/http: request canceled (Client.Timeout exceeded ...) error.
	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	return httpClient
}
