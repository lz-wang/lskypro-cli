package cmd

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/lzwang/lskypro-cli/internal/output"
	"github.com/urfave/cli/v2"
)

var uploadCommand = &cli.Command{
	Name:    "upload",
	Aliases: []string{"up"},
	Usage:   "上传图片",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "strategy",
			Aliases: []string{"s"},
			Usage:   "存储策略 ID",
		},
		&cli.StringFlag{
			Name:  "copy",
			Value: "url",
			Usage: "复制链接格式: url/markdown/bbcode/html",
		},
	},
	Action: doUpload,
}

func doUpload(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 获取文件路径
	if cCtx.NArg() == 0 {
		return fmt.Errorf("请指定要上传的图片文件")
	}

	filePath := cCtx.Args().First()
	strategyID := cCtx.Int("strategy")
	copyFormat := cCtx.String("copy")

	// 调用 API 上传图片
	result, err := apiClient.UploadImage(filePath, strategyID)
	if err != nil {
		return fmt.Errorf("上传失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("上传失败: %s", result.Message)
	}

	// 格式化输出
	data := &output.UploadData{
		Key:        result.Data.Key,
		OriginName: result.Data.OriginName,
		Size:       result.Data.Size,
		Mimetype:   result.Data.Mimetype,
		Width:      result.Data.Width,
		Height:     result.Data.Height,
		MD5:        result.Data.MD5,
		SHA1:       result.Data.SHA1,
		Links: output.Links{
			URL:              result.Data.Links.URL,
			HTML:             result.Data.Links.HTML,
			BBCode:           result.Data.Links.BBCode,
			Markdown:         result.Data.Links.Markdown,
			MarkdownWithLink: result.Data.Links.MarkdownWithLink,
			ThumbnailURL:     result.Data.Links.ThumbnailURL,
		},
	}

	formatter.FormatUpload(data)

	// 根据指定格式输出链接到 stderr (方便复制)
	var copyLink string
	switch copyFormat {
	case "markdown":
		copyLink = result.Data.Links.Markdown
	case "bbcode":
		copyLink = result.Data.Links.BBCode
	case "html":
		copyLink = result.Data.Links.HTML
	default:
		copyLink = result.Data.Links.URL
	}
	// 复制到剪贴板
	if err := clipboard.WriteAll(copyLink); err != nil {
		fmt.Fprintf(os.Stderr, "警告: 复制到剪贴板失败: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "已复制: %s\n", copyLink)
	}

	return nil
}
