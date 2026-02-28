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

var albumsCommand = &cli.Command{
	Name:    "albums",
	Aliases: []string{"alb"},
	Usage:   "管理相册",
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "列出相册",
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
					Usage: "排序: newest/earliest/most/least",
				},
				&cli.StringFlag{
					Name:    "keyword",
					Aliases: []string{"k"},
					Usage:   "搜索关键词",
				},
			},
			Action: doAlbumsList,
		},
		{
			Name:    "delete",
			Aliases: []string{"del", "rm"},
			Usage:   "删除相册",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{"f"},
					Usage:   "跳过确认",
				},
			},
			Action: doAlbumsDelete,
		},
	},
}

func doAlbumsList(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 构建请求选项
	opts := &api.ListAlbumsOptions{
		Page:    cCtx.Int("page"),
		Order:   cCtx.String("order"),
		Keyword: cCtx.String("keyword"),
	}

	// 调用 API 获取相册列表
	result, err := apiClient.ListAlbums(opts)
	if err != nil {
		return fmt.Errorf("获取相册列表失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("获取相册列表失败: %s", result.Message)
	}

	// 转换数据
	albums := make([]output.AlbumData, 0, len(result.Data.Albums))
	for _, album := range result.Data.Albums {
		albums = append(albums, output.AlbumData{
			ID:       album.ID,
			Name:     album.Name,
			Intro:    album.Intro,
			ImageNum: album.ImageNum,
		})
	}

	data := &output.AlbumsListData{
		Albums: albums,
	}

	formatter.FormatAlbums(data)
	return nil
}

func doAlbumsDelete(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 获取相册 ID
	if cCtx.NArg() == 0 {
		return fmt.Errorf("请指定要删除的相册 ID")
	}

	ids := cCtx.Args().Slice()
	force := cCtx.Bool("force")

	// 确认删除
	if !force {
		fmt.Printf("即将删除 %d 个相册:\n", len(ids))
		for _, id := range ids {
			fmt.Printf("  - %s\n", id)
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
	for _, idStr := range ids {
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			fmt.Printf("无效的相册 ID: %s\n", idStr)
			continue
		}

		if err := apiClient.DeleteAlbum(id); err != nil {
			fmt.Printf("删除失败 [%d]: %v\n", id, err)
		} else {
			fmt.Printf("已删除: %d\n", id)
		}
	}

	return nil
}
