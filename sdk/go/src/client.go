package megaverse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	token      string
	httpClient *http.Client
}

func NewClient(baseURL string, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+path, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Set("X-API-Key", c.apiKey)
	}
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// Auth
func (c *Client) Register(email, name, password string) (*User, error) {
	body := map[string]string{"email": email, "name": name, "password": password}
	data, err := c.doRequest("POST", "/auth/register", body)
	if err != nil {
		return nil, err
	}
	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) Login(email, password string) (*AuthResponse, error) {
	body := map[string]string{"email": email, "password": password}
	data, err := c.doRequest("POST", "/auth/login", body)
	if err != nil {
		return nil, err
	}
	var auth AuthResponse
	if err := json.Unmarshal(data, &auth); err != nil {
		return nil, err
	}
	c.SetToken(auth.AccessToken)
	return &auth, nil
}

// Users
func (c *Client) GetUser(id string) (*User, error) {
	data, err := c.doRequest("GET", "/users/"+id, nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) Follow(userID string) error {
	_, err := c.doRequest("POST", "/users/"+userID+"/follow", nil)
	return err
}

func (c *Client) Unfollow(userID string) error {
	_, err := c.doRequest("POST", "/users/"+userID+"/unfollow", nil)
	return err
}

// Posts
func (c *Client) CreatePost(content string) (*Post, error) {
	body := map[string]string{"content": content}
	data, err := c.doRequest("POST", "/posts", body)
	if err != nil {
		return nil, err
	}
	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, err
	}
	return &post, nil
}

func (c *Client) GetPost(id string) (*Post, error) {
	data, err := c.doRequest("GET", "/posts/"+id, nil)
	if err != nil {
		return nil, err
	}
	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, err
	}
	return &post, nil
}

func (c *Client) DeletePost(id string) error {
	_, err := c.doRequest("DELETE", "/posts/"+id, nil)
	return err
}

func (c *Client) GetFeed(params *PaginationParams) ([]Post, error) {
	path := "/feed"
	if params != nil {
		path = fmt.Sprintf("/feed?page=%d&limit=%d", params.Page, params.Limit)
	}
	data, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Posts []Post `json:"posts"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result.Posts, nil
}

// Health
func (c *Client) Health() (map[string]string, error) {
	data, err := c.doRequest("GET", "/health", nil)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
