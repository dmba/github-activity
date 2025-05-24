package github

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	DefaultBaseURL    = "https://api.github.com"
	DefaultApiVersion = "2022-11-28"
	DefaultTimeout    = time.Second * 30
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIClient struct {
	*http.Client
	BaseURL    string
	ApiVersion string
	Token      string
}

func NewGitHubAPIClient(token string) *APIClient {
	return &APIClient{
		BaseURL:    DefaultBaseURL,
		ApiVersion: DefaultApiVersion,
		Token:      token,
		Client: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

type PageOptions struct {
	Page    int
	PerPage int
}

func (c *APIClient) GetUsersEvents(ctx context.Context, username string, options *PageOptions) (*[]Event, error) {
	options = cmp.Or(options, &PageOptions{
		Page:    1,
		PerPage: 100,
	})

	endpoint := fmt.Sprintf("%s/users/%s/events?page=%d&per_page=%d", c.BaseURL, username, options.Page, options.PerPage)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	var res []Event
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *APIClient) sendRequest(req *http.Request, v any) error {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", c.ApiVersion)
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	if err = res.Body.Close(); err != nil {
		return err
	}

	return nil
}
