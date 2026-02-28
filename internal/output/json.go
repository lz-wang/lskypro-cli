package output

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONFormatter JSON 格式化器
type JSONFormatter struct {
	writer io.Writer
}

// NewJSONFormatter 创建 JSON 格式化器
func NewJSONFormatter(writer io.Writer) *JSONFormatter {
	return &JSONFormatter{writer: writer}
}

// FormatProfile 格式化用户资料
func (f *JSONFormatter) FormatProfile(data *ProfileData) string {
	output, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(f.writer, string(output))
	return ""
}

// FormatUpload 格式化上传结果
func (f *JSONFormatter) FormatUpload(data *UploadData) string {
	output, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(f.writer, string(output))
	return ""
}

// FormatImages 格式化图片列表
func (f *JSONFormatter) FormatImages(data *ImagesListData) string {
	output, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(f.writer, string(output))
	return ""
}

// FormatAlbums 格式化相册列表
func (f *JSONFormatter) FormatAlbums(data *AlbumsListData) string {
	output, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(f.writer, string(output))
	return ""
}

// FormatStrategies 格式化存储策略列表
func (f *JSONFormatter) FormatStrategies(data *StrategiesListData) string {
	output, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(f.writer, string(output))
	return ""
}
