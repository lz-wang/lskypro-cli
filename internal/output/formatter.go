package output

import (
	"fmt"
	"io"
)

// Formatter 输出格式化接口
type Formatter interface {
	FormatProfile(data *ProfileData) string
	FormatUpload(data *UploadData) string
	FormatImages(data *ImagesListData) string
	FormatAlbums(data *AlbumsListData) string
	FormatStrategies(data *StrategiesListData) string
}

// ProfileData 用户资料数据
type ProfileData struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	StorageUsed  int64  `json:"storage_used"`
	ImageNum     int    `json:"image_num"`
	AlbumNum     int    `json:"album_num"`
	RegisteredIP string `json:"registered_ip"`
}

// UploadData 上传结果数据
type UploadData struct {
	Key        string `json:"key"`
	OriginName string `json:"origin_name"`
	Size       float64 `json:"size"`
	Mimetype   string `json:"mimetype"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	MD5        string `json:"md5"`
	SHA1       string `json:"sha1"`
	Links      Links  `json:"links"`
}

// Links 链接数据
type Links struct {
	URL              string `json:"url"`
	HTML             string `json:"html"`
	BBCode           string `json:"bbcode"`
	Markdown         string `json:"markdown"`
	MarkdownWithLink string `json:"markdown_with_link"`
	ThumbnailURL     string `json:"thumbnail_url"`
}

// ImageData 图片数据
type ImageData struct {
	Key        string `json:"key"`
	OriginName string `json:"origin_name"`
	Size       float64 `json:"size"`
	Mimetype   string `json:"mimetype"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	MD5        string `json:"md5"`
	SHA1       string `json:"sha1"`
	HumanDate  string `json:"human_date"`
	Date       string `json:"date"`
	Links      Links  `json:"links"`
}

// ImagesListData 图片列表数据
type ImagesListData struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
	Images      []ImageData `json:"images"`
}

// AlbumData 相册数据
type AlbumData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Intro    string `json:"intro"`
	ImageNum int    `json:"image_num"`
}

// AlbumsListData 相册列表数据
type AlbumsListData struct {
	Albums []AlbumData `json:"albums"`
}

// StrategyData 存储策略数据
type StrategyData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// StrategiesListData 存储策略列表数据
type StrategiesListData struct {
	Strategies []StrategyData `json:"strategies"`
}

// FormatType 格式类型
type FormatType string

const (
	FormatTable FormatType = "table"
	FormatJSON  FormatType = "json"
	FormatPlain FormatType = "plain"
)

// NewFormatter 创建格式化器
func NewFormatter(format string, writer io.Writer) Formatter {
	switch FormatType(format) {
	case FormatJSON:
		return &JSONFormatter{writer: writer}
	case FormatPlain:
		return &PlainFormatter{writer: writer}
	default:
		return &TableFormatter{writer: writer}
	}
}

// FormatBytes 格式化字节数
func FormatBytes(bytes float64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%.0f B", bytes)
	}
	div, exp := float64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", bytes/div, "KMGTPE"[exp])
}
