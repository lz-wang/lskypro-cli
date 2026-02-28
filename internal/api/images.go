package api

import (
	"fmt"
	"io"
	"os"
)

// UploadResponse 上传响应
type UploadResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Key        string `json:"key"`
		OriginName string `json:"origin_name"`
		Size       int64  `json:"size"`
		Mimetype   string `json:"mimetype"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		MD5        string `json:"md5"`
		SHA1       string `json:"sha1"`
		Links      struct {
			URL              string `json:"url"`
			HTML             string `json:"html"`
			BBCode           string `json:"bbcode"`
			Markdown         string `json:"markdown"`
			MarkdownWithLink string `json:"markdown_with_link"`
			ThumbnailURL     string `json:"thumbnail_url"`
		} `json:"links"`
	} `json:"data"`
}

// ListImagesOptions 图片列表选项
type ListImagesOptions struct {
	Page       int    `json:"page"`
	Order      string `json:"order"`
	Permission string `json:"permission"`
	AlbumID    int    `json:"album_id"`
	Keyword    string `json:"keyword"`
}

// ImagesListResponse 图片列表响应
type ImagesListResponse struct {
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
		Images      []struct {
			Key        string `json:"key"`
			OriginName string `json:"origin_name"`
			Size       int64  `json:"size"`
			Mimetype   string `json:"mimetype"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			MD5        string `json:"md5"`
			SHA1       string `json:"sha1"`
			HumanDate  string `json:"human_date"`
			Date       string `json:"date"`
			Links      struct {
				URL              string `json:"url"`
				HTML             string `json:"html"`
				BBCode           string `json:"bbcode"`
				Markdown         string `json:"markdown"`
				MarkdownWithLink string `json:"markdown_with_link"`
				ThumbnailURL     string `json:"thumbnail_url"`
			} `json:"links"`
		} `json:"images"`
	} `json:"data"`
}

// UploadImage 上传图片
func (c *Client) UploadImage(filePath string, strategyID int) (*UploadResponse, error) {
	var result UploadResponse

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, &APIError{
			StatusCode: 0,
			Message:    "文件不存在: " + filePath,
		}
	}

	req := c.httpClient.R().
		SetAuthToken(c.token).
		SetFile("file", filePath).
		SetResult(&result)

	if strategyID > 0 {
		req.SetFormData(map[string]string{
			"strategy_id": fmt.Sprintf("%d", strategyID),
		})
	}

	resp, err := req.Post("/upload")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

// UploadImageFromReader 从 Reader 上传图片
func (c *Client) UploadImageFromReader(reader io.Reader, filename string, strategyID int) (*UploadResponse, error) {
	var result UploadResponse

	req := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result)

	// 使用 SetMultipartField 设置 multipart 字段
	req.SetMultipartField("file", filename, "", reader)

	if strategyID > 0 {
		req.SetFormData(map[string]string{
			"strategy_id": fmt.Sprintf("%d", strategyID),
		})
	}

	resp, err := req.Post("/upload")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListImages 获取图片列表
func (c *Client) ListImages(opts *ListImagesOptions) (*ImagesListResponse, error) {
	var result ImagesListResponse

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
		if opts.Permission != "" {
			params["permission"] = opts.Permission
		}
		if opts.AlbumID > 0 {
			params["album_id"] = fmt.Sprintf("%d", opts.AlbumID)
		}
		if opts.Keyword != "" {
			params["keyword"] = opts.Keyword
		}
		if len(params) > 0 {
			req.SetQueryParams(params)
		}
	}

	resp, err := req.Get("/images")

	if err := checkResponse(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteImage 删除图片
func (c *Client) DeleteImage(key string) error {
	var result BaseResponse

	resp, err := c.httpClient.R().
		SetAuthToken(c.token).
		SetResult(&result).
		SetPathParams(map[string]string{
			"key": key,
		}).
		Delete("/images/{key}")

	if err := checkResponse(resp, err); err != nil {
		return err
	}

	return nil
}
