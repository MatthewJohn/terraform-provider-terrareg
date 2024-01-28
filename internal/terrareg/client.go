package terrareg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TerraregClient struct {
	Url    string
	ApiKey string
}

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

func (c *TerraregClient) CreateNamespace(name string) (*http.Response, error) {

	url := c.getTerraregApiUrl("namespaces")

	postData := map[string]string{
		"name": name,
	}
	res, err := c.makePostRequest(url, postData)
	if err != nil {
		return nil, err
	}

	return res, nil
}
