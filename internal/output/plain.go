package output

import (
	"fmt"
	"io"
)

// PlainFormatter 纯文本格式化器
type PlainFormatter struct {
	writer io.Writer
}

// NewPlainFormatter 创建纯文本格式化器
func NewPlainFormatter(writer io.Writer) *PlainFormatter {
	return &PlainFormatter{writer: writer}
}

// FormatProfile 格式化用户资料
func (f *PlainFormatter) FormatProfile(data *ProfileData) string {
	fmt.Fprintf(f.writer, "ID: %d\n", data.ID)
	fmt.Fprintf(f.writer, "邮箱: %s\n", data.Email)
	fmt.Fprintf(f.writer, "用户名: %s\n", data.Name)
	fmt.Fprintf(f.writer, "头像: %s\n", data.Avatar)
	fmt.Fprintf(f.writer, "已用空间: %s\n", FormatBytes(float64(data.StorageUsed)))
	fmt.Fprintf(f.writer, "图片数量: %d\n", data.ImageNum)
	fmt.Fprintf(f.writer, "相册数量: %d\n", data.AlbumNum)
	fmt.Fprintf(f.writer, "注册 IP: %s\n", data.RegisteredIP)
	return ""
}

// FormatUpload 格式化上传结果
func (f *PlainFormatter) FormatUpload(data *UploadData) string {
	fmt.Fprintf(f.writer, "Key: %s\n", data.Key)
	fmt.Fprintf(f.writer, "原文件名: %s\n", data.OriginName)
	fmt.Fprintf(f.writer, "大小: %s\n", FormatBytes(data.Size))
	fmt.Fprintf(f.writer, "类型: %s\n", data.Mimetype)
	fmt.Fprintf(f.writer, "尺寸: %d x %d\n", data.Width, data.Height)
	fmt.Fprintf(f.writer, "MD5: %s\n", data.MD5)
	fmt.Fprintf(f.writer, "SHA1: %s\n", data.SHA1)
	fmt.Fprintf(f.writer, "链接: %s\n", data.Links.URL)
	fmt.Fprintf(f.writer, "Markdown: %s\n", data.Links.Markdown)
	fmt.Fprintf(f.writer, "HTML: %s\n", data.Links.HTML)
	fmt.Fprintf(f.writer, "BBCode: %s\n", data.Links.BBCode)
	fmt.Fprintf(f.writer, "缩略图: %s\n", data.Links.ThumbnailURL)
	return ""
}

// FormatImages 格式化图片列表
func (f *PlainFormatter) FormatImages(data *ImagesListData) string {
	fmt.Fprintf(f.writer, "共 %d 张图片，第 %d/%d 页\n\n", data.Total, data.CurrentPage, data.LastPage)

	for _, img := range data.Images {
		fmt.Fprintf(f.writer, "%s\t%s\t%s\t%dx%d\t%s\n",
			img.Key,
			img.OriginName,
			FormatBytes(img.Size),
			img.Width,
			img.Height,
			img.HumanDate,
		)
	}
	return ""
}

// FormatAlbums 格式化相册列表
func (f *PlainFormatter) FormatAlbums(data *AlbumsListData) string {
	for _, album := range data.Albums {
		fmt.Fprintf(f.writer, "%d\t%s\t%s\t%d\n",
			album.ID,
			album.Name,
			album.Intro,
			album.ImageNum,
		)
	}
	return ""
}

// FormatStrategies 格式化存储策略列表
func (f *PlainFormatter) FormatStrategies(data *StrategiesListData) string {
	for _, strategy := range data.Strategies {
		fmt.Fprintf(f.writer, "%d\t%s\n", strategy.ID, strategy.Name)
	}
	return ""
}
