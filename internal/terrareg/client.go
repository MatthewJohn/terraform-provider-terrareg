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

type NamespaceModel struct {
	DisplayName    string `json:"display_name"`
	IsAutoVerified bool   `json:"is_auto_verified"`
	Trusted        bool   `json:"trusted"`
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

func (c *TerraregClient) makeRequest(url string, requestMethod string, jsonData any) (*http.Response, error) {
	jsonStr := "{}"
	if jsonData != nil {
		var err error = nil
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
		jsonStr = string(jsonBytes)
	}

	req, err := http.NewRequest(requestMethod, url, strings.NewReader(jsonStr))
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

func (c *TerraregClient) CreateNamespace(name string) error {

	url := c.getTerraregApiUrl("namespaces")

	postData := map[string]string{
		"name": name,
	}
	res, err := c.makeRequest(url, "POST", postData)
	if err != nil {
		return err
	}

	err = c.handleCommonStatusCode(res.StatusCode)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return ErrUnknownError
	}
	return nil
}

func (c *TerraregClient) GetNamespace(name string) (*NamespaceModel, error) {
	url := c.getTerraregApiUrl(fmt.Sprintf("namespaces/%s", name))

	res, err := c.makeRequest(url, "GET", nil)
	if err != nil {
		return nil, err
	}

	err = c.handleCommonStatusCode(res.StatusCode)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, ErrUnknownError
	}

	// Body is 200
	if res.Body == nil {
		return nil, ErrUnknownError
	}

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()

	var namespace NamespaceModel
	err = dec.Decode(&namespace)
	if err != nil {
		fmt.Printf("Terrareg Client: Unable to decode namespace JSON from response body")
		return nil, err
	}
	return &namespace, nil
}

func (c *TerraregClient) DeleteNamespace(name string) error {
	url := c.getTerraregApiUrl(fmt.Sprintf("namespaces/%s", name))

	res, err := c.makeRequest(url, "DELETE", nil)
	if err != nil {
		return err
	}

	err = c.handleCommonStatusCode(res.StatusCode)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return ErrUnknownError
	}

	return nil
}
