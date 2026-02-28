package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lzwang/lskypro-cli/internal/api"
	"github.com/lzwang/lskypro-cli/internal/output"
	"github.com/urfave/cli/v2"
)

var imagesCommand = &cli.Command{
	Name:    "images",
	Aliases: []string{"img"},
	Usage:   "管理图片",
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "列出图片",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "page",
					Aliases: []string{"p"},
					Value: 1,
					Usage: "页码",
				},
				&cli.StringFlag{
					Name:  "order",
					Value: "newest",
					Usage: "排序: newest/earliest/utmost/least",
				},
				&cli.StringFlag{
					Name:  "permission",
					Usage: "权限: public/private",
				},
				&cli.IntFlag{
					Name:  "album",
					Aliases: []string{"a"},
					Usage: "相册 ID",
				},
				&cli.StringFlag{
					Name:    "keyword",
					Aliases: []string{"k"},
					Usage:   "搜索关键词",
				},
			},
			Action: doImagesList,
		},
		{
			Name:    "delete",
			Aliases: []string{"del", "rm"},
			Usage:   "删除图片",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{"f"},
					Usage:   "跳过确认",
				},
			},
			Action: doImagesDelete,
		},
	},
}

func doImagesList(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 构建请求选项
	opts := &api.ListImagesOptions{
		Page:       cCtx.Int("page"),
		Order:      cCtx.String("order"),
		Permission: cCtx.String("permission"),
		AlbumID:    cCtx.Int("album"),
		Keyword:    cCtx.String("keyword"),
	}

	// 调用 API 获取图片列表
	result, err := apiClient.ListImages(opts)
	if err != nil {
		return fmt.Errorf("获取图片列表失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("获取图片列表失败: %s", result.Message)
	}

	// 转换数据
	images := make([]output.ImageData, 0, len(result.Data.Images))
	for _, img := range result.Data.Images {
		images = append(images, output.ImageData{
			Key:        img.Key,
			OriginName: img.OriginName,
			Size:       img.Size,
			Mimetype:   img.Mimetype,
			Width:      img.Width,
			Height:     img.Height,
			MD5:        img.MD5,
			SHA1:       img.SHA1,
			HumanDate:  img.HumanDate,
			Date:       img.Date,
			Links: output.Links{
				URL:              img.Links.URL,
				HTML:             img.Links.HTML,
				BBCode:           img.Links.BBCode,
				Markdown:         img.Links.Markdown,
				MarkdownWithLink: img.Links.MarkdownWithLink,
				ThumbnailURL:     img.Links.ThumbnailURL,
			},
		})
	}

	data := &output.ImagesListData{
		CurrentPage: result.Data.CurrentPage,
		LastPage:    result.Data.LastPage,
		PerPage:     result.Data.PerPage,
		Total:       result.Data.Total,
		Images:      images,
	}

	formatter.FormatImages(data)
	return nil
}

func doImagesDelete(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 获取图片 Key
	if cCtx.NArg() == 0 {
		return fmt.Errorf("请指定要删除的图片 Key")
	}

	keys := cCtx.Args().Slice()
	force := cCtx.Bool("force")

	// 确认删除
	if !force {
		fmt.Printf("即将删除 %d 张图片:\n", len(keys))
		for _, key := range keys {
			fmt.Printf("  - %s\n", key)
		}
		fmt.Print("确认删除? (y/N): ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" && input != "yes" {
			fmt.Println("已取消")
			return nil
		}
	}

	// 逐个删除
	for _, key := range keys {
		if err := apiClient.DeleteImage(key); err != nil {
			fmt.Printf("删除失败 [%s]: %v\n", key, err)
		} else {
			fmt.Printf("已删除: %s\n", key)
		}
	}

	return nil
}
