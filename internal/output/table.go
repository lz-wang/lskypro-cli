package output

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

// TableFormatter 表格格式化器
type TableFormatter struct {
	writer io.Writer
}

// NewTableFormatter 创建表格格式化器
func NewTableFormatter(writer io.Writer) *TableFormatter {
	return &TableFormatter{writer: writer}
}

func (f *TableFormatter) createTable() *tablewriter.Table {
	table := tablewriter.NewTable(f.writer,
		tablewriter.WithBorders(tw.Border{Left: tw.Off, Right: tw.Off, Top: tw.Off, Bottom: tw.Off}),
	)
	return table
}

// FormatProfile 格式化用户资料
func (f *TableFormatter) FormatProfile(data *ProfileData) string {
	table := f.createTable()

	rows := [][]string{
		{"ID", fmt.Sprintf("%d", data.ID)},
		{"邮箱", data.Email},
		{"用户名", data.Name},
		{"头像", data.Avatar},
		{"已用空间", FormatBytes(data.StorageUsed)},
		{"图片数量", fmt.Sprintf("%d", data.ImageNum)},
		{"相册数量", fmt.Sprintf("%d", data.AlbumNum)},
		{"注册 IP", data.RegisteredIP},
	}

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()
	return ""
}

// FormatUpload 格式化上传结果
func (f *TableFormatter) FormatUpload(data *UploadData) string {
	table := f.createTable()

	rows := [][]string{
		{"Key", data.Key},
		{"原文件名", data.OriginName},
		{"大小", FormatBytes(data.Size)},
		{"类型", data.Mimetype},
		{"尺寸", fmt.Sprintf("%d x %d", data.Width, data.Height)},
		{"MD5", data.MD5},
		{"SHA1", data.SHA1},
		{"链接", data.Links.URL},
		{"Markdown", data.Links.Markdown},
		{"HTML", data.Links.HTML},
		{"BBCode", data.Links.BBCode},
		{"缩略图", data.Links.ThumbnailURL},
	}

	for _, row := range rows {
		table.Append(row)
	}

	table.Render()

	// 输出到 stderr，方便复制链接
	fmt.Fprintf(os.Stderr, "\n链接: %s\n", data.Links.URL)
	return ""
}

// FormatImages 格式化图片列表
func (f *TableFormatter) FormatImages(data *ImagesListData) string {
	fmt.Fprintf(f.writer, "共 %d 张图片，第 %d/%d 页\n\n", data.Total, data.CurrentPage, data.LastPage)

	if len(data.Images) == 0 {
		fmt.Fprintln(f.writer, "暂无图片")
		return ""
	}

	table := f.createTable()

	// 设置表头
	table.Header([]string{"Key", "原文件名", "大小", "尺寸", "上传时间"})

	for _, img := range data.Images {
		table.Append([]string{
			img.Key,
			img.OriginName,
			FormatBytes(img.Size),
			fmt.Sprintf("%dx%d", img.Width, img.Height),
			img.HumanDate,
		})
	}

	table.Render()
	return ""
}

// FormatAlbums 格式化相册列表
func (f *TableFormatter) FormatAlbums(data *AlbumsListData) string {
	if len(data.Albums) == 0 {
		fmt.Fprintln(f.writer, "暂无相册")
		return ""
	}

	table := f.createTable()

	table.Header([]string{"ID", "名称", "简介", "图片数量"})

	for _, album := range data.Albums {
		table.Append([]string{
			fmt.Sprintf("%d", album.ID),
			album.Name,
			album.Intro,
			fmt.Sprintf("%d", album.ImageNum),
		})
	}

	table.Render()
	return ""
}

// FormatStrategies 格式化存储策略列表
func (f *TableFormatter) FormatStrategies(data *StrategiesListData) string {
	if len(data.Strategies) == 0 {
		fmt.Fprintln(f.writer, "暂无存储策略")
		return ""
	}

	table := f.createTable()

	table.Header([]string{"ID", "名称"})

	for _, strategy := range data.Strategies {
		table.Append([]string{
			fmt.Sprintf("%d", strategy.ID),
			strategy.Name,
		})
	}

	table.Render()
	return ""
}
