package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// The Client provides access to the notion resources.
type Client interface {
	Search(ctx context.Context, searchRequest SearchRequest) (searchResponse SearchResponse, err error)
}

type client struct {
	httpClient *http.Client
}

func (c *client) fetch(ctx context.Context, method string, path string, reqValue any, resValue any) error {
	body, err := json.Marshal(reqValue)
	if err != nil {
		return fmt.Errorf("failed to marshal the request body: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		"https://api.notion.com/v1"+path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return fmt.Errorf("failed to init request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_API_KEY"))
	req.Header.Add("Notion-Version", "2022-06-28")

	response, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call request: %w", err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(resValue)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	return nil
}

func (c *client) Search(ctx context.Context, searchRequest SearchRequest) (SearchResponse, error) {
	var searchResponse SearchResponse

	err := c.fetch(ctx, http.MethodPost, "/search", &searchRequest, &searchResponse)
	if err != nil {
		return searchResponse, err
	}

	return searchResponse, nil
}

// New returns a new instance of notion.Client.
func New() Client {
	return &client{
		httpClient: &http.Client{},
	}
}
