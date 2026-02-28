package api

// StrategiesListResponse 存储策略列表响应
type StrategiesListResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Strategies []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"strategies"`
	} `json:"data"`
}

// ListStrategies 获取存储策略列表
func (c *Client) ListStrategies(keyword string) (*StrategiesListResponse, error) {
	var result StrategiesListResponse

	req := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result)

	if keyword != "" {
		req.SetQueryParam("keyword", keyword)
	}

	resp, err := req.Get("/strategies")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
