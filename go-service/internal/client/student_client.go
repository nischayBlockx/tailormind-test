package client

import (
	"encoding/json"
	"fmt"
	"go-service/internal/models"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) GetStudentByIDHttp(id string) (*models.StudentDetails, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", c.BaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call student service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("student service returned status: %d", resp.StatusCode)
	}

	var result struct {
		Data models.StudentDetails `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result.Data, nil
}
