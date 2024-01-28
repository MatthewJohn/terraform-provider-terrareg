package terrareg

import (
	"encoding/json"
	"fmt"
)

type NamespaceModel struct {
	DisplayName    string `json:"display_name"`
	IsAutoVerified bool   `json:"is_auto_verified"`
	Trusted        bool   `json:"trusted"`
}

type NamespaceConfigModel struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func (c *TerraregClient) CreateNamespace(config NamespaceConfigModel) error {

	url := c.getTerraregApiUrl("namespaces")

	res, err := c.makeRequest(url, "POST", config)
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

func (c *TerraregClient) UpdateNamespace(name string, config NamespaceConfigModel) error {
	url := c.getTerraregApiUrl(fmt.Sprintf("namespaces/%s", name))

	res, err := c.makeRequest(url, "POST", config)
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
