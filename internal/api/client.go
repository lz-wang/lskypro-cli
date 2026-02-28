package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client API 客户端
type Client struct {
	baseURL    string
	token      string
	httpClient *resty.Client
}

// NewClient 创建 API 客户端
func NewClient(baseURL, token string) *Client {
	client := resty.New()
	client.SetHeader("Accept", "application/json")

	if baseURL != "" {
		client.SetBaseURL(baseURL + "/api/v1")
	}

	return &Client{
		baseURL:    baseURL,
		token:      token,
		httpClient: client,
	}
}

// SetBaseURL 设置服务器地址
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
	c.httpClient.SetBaseURL(url + "/api/v1")
}

// SetToken 设置认证 Token
func (c *Client) SetToken(token string) {
	c.token = token
	c.httpClient.SetAuthToken(token)
}

// GetToken 获取当前 Token
func (c *Client) GetToken() string {
	return c.token
}

// GetBaseURL 获取服务器地址
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// BaseResponse 基础响应结构
type BaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// APIError API 错误
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

// checkResponse 检查响应
func checkResponse(resp *resty.Response, err error) error {
	if err != nil {
		return &APIError{
			StatusCode: 0,
			Message:    fmt.Sprintf("请求失败: %v", err),
		}
	}

	if resp.IsError() {
		return &APIError{
			StatusCode: resp.StatusCode(),
			Message:    fmt.Sprintf("HTTP 错误: %s", resp.Status()),
		}
	}

	return nil
}
