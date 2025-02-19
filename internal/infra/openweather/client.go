package openweather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type VpBackendClient struct {
	client *http.Client
	url    string
}

func NewOpenWeatherClient(conf Config) *VpBackendClient {
	var url string
	if conf.Port != 0 {
		url = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	} else {
		url = conf.Host
	}
	return &VpBackendClient{
		client: &http.Client{Timeout: 30 * time.Second},
		url:    url,
	}
}

func (c *VpBackendClient) getRequest(route string, queryParams map[string]string) ([]byte, error) {
	requestURL := c.url + route

	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request URL: %v", err)
	}

	if queryParams != nil {
		query := u.Query()
		for key, value := range queryParams {
			query.Add(key, value)
		}
		u.RawQuery = query.Encode()
	}

	resp, err := c.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request %s failed with status: %s", u.String(), resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
