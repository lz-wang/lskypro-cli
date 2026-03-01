package api

import "fmt"

// ListAlbumsOptions 相册列表选项
type ListAlbumsOptions struct {
	Page    int    `json:"page"`
	Order   string `json:"order"`
	Keyword string `json:"keyword"`
}

// AlbumsListResponse 相册列表响应
type AlbumsListResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CurrentPage int `json:"current_page"`
		FirstPage   int `json:"first_page"`
		From        int `json:"from"`
		LastPage    int `json:"last_page"`
		PerPage     int `json:"per_page"`
		To          int `json:"to"`
		Total       int `json:"total"`
		Albums      []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Intro    string `json:"intro"`
			ImageNum int    `json:"image_num"`
		} `json:"data"`
	} `json:"data"`
}

// ListAlbums 获取相册列表
func (c *Client) ListAlbums(opts *ListAlbumsOptions) (*AlbumsListResponse, error) {
	var result AlbumsListResponse

	req := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result)

	if opts != nil {
		params := make(map[string]string)
		if opts.Page > 0 {
			params["page"] = fmt.Sprintf("%d", opts.Page)
		}
		if opts.Order != "" {
			params["order"] = opts.Order
		}
		if opts.Keyword != "" {
			params["keyword"] = opts.Keyword
		}
		if len(params) > 0 {
			req.SetQueryParams(params)
		}
	}

	resp, err := req.Get("/albums")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAlbum 删除相册
func (c *Client) DeleteAlbum(id int) error {
	var result BaseResponse

	resp, err := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result).
		SetPathParams(map[string]string{
			"id": fmt.Sprintf("%d", id),
		}).
		Delete("/albums/{id}")

	if err := checkResponse(resp, err); err != nil {
		return err
	}

	return nil
}
