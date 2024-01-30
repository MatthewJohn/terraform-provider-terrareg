package terrareg

import (
	"encoding/json"
	"fmt"
)

type GitProviderModel struct {
	ID   int    `json:"id" tfsdk:"id"`
	Name string `json:"name" tfsdk:"name"`
}

func (c *TerraregClient) GetGitProviders() ([]GitProviderModel, error) {
	url := c.getTerraregApiUrl("git_providers")

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
	// dec.DisallowUnknownFields()

	var data []GitProviderModel
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("Terrareg Client: Unable to decode git providers JSON from response body")
		return nil, err
	}
	return data, nil
}
