package terrareg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type TerraregClient struct {
	Url    string
	ApiKey string
}

var ErrNotFound = errors.New("Not found")
var ErrInvalidAuth = errors.New("Invalid Authentication")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrUnknownServerError = errors.New("Unknown Server error")
var ErrUnknownError = errors.New("Unknown Error")

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

func (c *TerraregClient) makePostRequest(url string, jsonData any) (*http.Response, error) {
	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonStr)))
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

func (c *TerraregClient) CreateNamespace(name string) error {

	url := c.getTerraregApiUrl("namespaces")

	postData := map[string]string{
		"name": name,
	}
	res, err := c.makePostRequest(url, postData)
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode == 401 {
		return ErrInvalidAuth
	}
	if res.StatusCode == 403 {
		return ErrUnauthorized
	}
	if res.StatusCode >= 500 && res.StatusCode <= 503 {
		return ErrUnknownServerError
	}
	return ErrUnknownError
}
