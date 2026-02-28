package api

// TokenRequest 登录请求
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TokenResponse Token 响应
type TokenResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

// ProfileResponse 用户资料响应
type ProfileResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Avatar       string `json:"avatar"`
		StorageUsed  int64  `json:"storage_used"`
		ImageNum     int    `json:"image_num"`
		AlbumNum     int    `json:"album_num"`
		RegisteredIP string `json:"registered_ip"`
	} `json:"data"`
}

// Login 登录获取 Token
func (c *Client) Login(email, password string) (*TokenResponse, error) {
	var result TokenResponse

	resp, err := c.httpClient.R().
		SetBody(TokenRequest{
			Email:    email,
			Password: password,
		}).
		SetResult(&result).
		Post("/tokens")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

// Logout 清空 Token
func (c *Client) Logout() error {
	var result BaseResponse

	resp, err := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result).
		Delete("/tokens")

	if err := checkResponse(resp, err); err != nil {
		return err
	}

	return nil
}

// GetProfile 获取用户资料
func (c *Client) GetProfile() (*ProfileResponse, error) {
	var result ProfileResponse

	resp, err := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result).
		Get("/profile")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
