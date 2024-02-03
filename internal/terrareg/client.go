package terrareg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type TerraregClient struct {
	Url    string
	ApiKey string
}

var ErrNotFound = errors.New("Not found")
var ErrInvalidAuth = errors.New("Invalid Authentication")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrUnknownServerError = errors.New("Unknown Server error")
var ErrUnknownError = errors.New("Unknown HTTP Response")

func NewClient(url string, apiKey string) (*TerraregClient, error) {
	return &TerraregClient{
		Url:    url,
		ApiKey: apiKey,
	}, nil
}

func (c *TerraregClient) getHeaders() http.Header {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")
	if c.ApiKey != "" {
		headers.Set("X-Terrareg-ApiKey", c.ApiKey)
	}

	return headers
}

func (c *TerraregClient) getHttpClient() http.Client {
	return http.Client{}
}

func (c *TerraregClient) getTerraregApiUrl(apiEndpoint string) string {
	return fmt.Sprintf("%s/v1/terrareg/%s", c.Url, apiEndpoint)
}

func (c *TerraregClient) printBody(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("[terrareg] Failed to dump repsonse")
		return
	}
	fmt.Printf("[terrareg] Got body repsonse: %s\n", string(respDump))
}

func (c *TerraregClient) makeRequest(url string, requestMethod string, jsonData any) (*http.Response, error) {
	body := new(bytes.Buffer)
	if jsonData != nil {
		err := json.NewEncoder(body).Encode(jsonData)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(requestMethod, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = c.getHeaders().Clone()
	httpClient := c.getHttpClient()

	httpRes, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return httpRes, nil
}

func (c *TerraregClient) handleCommonStatusCode(statusCode int) error {
	if statusCode == 401 {
		return ErrInvalidAuth
	}
	if statusCode == 403 {
		return ErrUnauthorized
	}
	if statusCode >= 500 && statusCode <= 503 {
		return ErrUnknownServerError
	}
	return nil
}
