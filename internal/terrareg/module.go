package terrareg

import (
	"encoding/json"
	"fmt"
)

type ModuleModel struct {
	GitProviderID         int64  `json:"git_provider_id"`
	RepoBaseUrlTemplate   string `json:"repo_base_url_template"`
	RepoCloneUrlTemplate  string `json:"repo_clone_url_template"`
	RepoBrowseUrlTemplate string `json:"repo_browse_url_template"`
	GitTagFormat          string `json:"git_tag_format"`
	GitPath               string `json:"git_path"`
}

type ModuleUpdateModel struct {
	*ModuleModel
	Namespace string `json:"namespace"`
	Name      string `json:"module"`
	Provider  string `json:"provider"`
}

func (c *TerraregClient) CreateModule(namespace string, name string, provider string, config ModuleModel) (string, error) {

	url := c.getTerraregApiUrl(fmt.Sprintf("modules/%s/%s/%s/create", namespace, name, provider))

	res, err := c.makeRequest(url, "POST", config)
	if err != nil {
		return "", err
	}

	err = c.handleCommonStatusCode(res.StatusCode)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", ErrUnknownError
	}

	if res.Body == nil {
		return "", ErrUnknownError
	}

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()

	type CreateRepsonse struct {
		ID string `json:"id"`
	}

	var data CreateRepsonse
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("Terrareg Client: Unable to decode module JSON from response body")
		return "", err
	}
	return data.ID, nil
}

func (c *TerraregClient) GetModule(namespace string, name string, provider string) (*ModuleModel, error) {
	url := c.getTerraregApiUrl(fmt.Sprintf("modules/%s/%s/%s", namespace, name, provider))

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

	var data ModuleModel
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("Terrareg Client: Unable to decode module JSON from response body")
		return nil, err
	}
	return &data, nil
}

func (c *TerraregClient) UpdateModule(namespace string, name string, provider string, config ModuleUpdateModel) (string, error) {

	url := c.getTerraregApiUrl(fmt.Sprintf("modules/%[1]s/%[2]s/%[3]s/settings", namespace, name, provider))

	// Ignore namespace/name/provider fields if they have not been set
	var dataToSend interface{}
	var newId string = ""
	if config.Namespace != "" && config.Name != "" && config.Provider != "" {
		dataToSend = config
		newId = fmt.Sprintf("%s/%s/%s", config.Namespace, config.Name, config.Provider)
	} else {
		dataToSend = config.ModuleModel
	}

	res, err := c.makeRequest(url, "POST", dataToSend)
	if err != nil {
		return "", err
	}

	err = c.handleCommonStatusCode(res.StatusCode)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", ErrUnknownError
	}

	// Terrareg doesn't provide an 'ID' response, at the moment,
	// so make one up on the fly
	return newId, nil
}

func (c *TerraregClient) DeleteModule(namespace string, name string, provider string) error {
	url := c.getTerraregApiUrl(fmt.Sprintf("modules/%s/%s/%s/delete", namespace, name, provider))

	// Since the namespace DELETE endpoint accepts JSON data,
	// an empty map must be passed to ensure the request is accepted.
	res, err := c.makeRequest(url, "DELETE", map[string]string{})
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
