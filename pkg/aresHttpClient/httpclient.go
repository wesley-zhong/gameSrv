package aresHttpClient

import (
	"bytes"
	"encoding/json"
	"gameSrv/pkg/log"
	"io"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second, // Total timeout for the request (dialing, handshaking, sending, receiving body)
		Transport: &http.Transport{
			MaxIdleConns:        100,              // Max idle connections across all hosts
			MaxIdleConnsPerHost: 10,               // Max idle connections to a single host
			IdleConnTimeout:     90 * time.Second, // How long an idle connection is kept alive
			// Add other transport configurations like TLSClientConfig, Proxy, DialContext etc.
		},
	}
}

func PostJson(url string, body interface{}, headers map[string]string) (resp []byte, err error) {
	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return postJsonByte(url, marshal, headers)

}
func postJsonByte(url string, body []byte, headers map[string]string) (resp []byte, err error) {
	return post(url, "application/json", body, headers)
}

func post(url string, contentType string, body []byte, headers map[string]string) (resp []byte, err error) {
	reqBody := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		log.Fatalf("Error creating POST request: %v", err)
	}
	req.Header.Set("Content-Type", contentType)
	// Add other headers if needed
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	postResp, err := httpClient.Do(req) // Use client.Do() for custom requests
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(postResp.Body)

	if postResp.StatusCode != http.StatusCreated {
		log.Fatalf("Unexpected POST status code: %d", postResp.StatusCode)
	}

	resp, err = io.ReadAll(postResp.Body)
	if err != nil {
		log.Fatalf("Error reading POST response body: %v", err)
	}
	return resp, nil
}
