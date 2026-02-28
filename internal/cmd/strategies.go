package cmd

import (
	"fmt"

	"github.com/lzwang/lskypro-cli/internal/output"
	"github.com/urfave/cli/v2"
)

var strategiesCommand = &cli.Command{
	Name:    "strategies",
	Aliases: []string{"st"},
	Usage:   "查看存储策略",
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "列出存储策略",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "keyword",
					Aliases: []string{"k"},
					Usage:   "搜索关键词",
				},
			},
			Action: doStrategiesList,
		},
	},
}

func doStrategiesList(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	keyword := cCtx.String("keyword")

	// 调用 API 获取存储策略列表
	result, err := apiClient.ListStrategies(keyword)
	if err != nil {
		return fmt.Errorf("获取存储策略列表失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("获取存储策略列表失败: %s", result.Message)
	}

	// 转换数据
	strategies := make([]output.StrategyData, 0, len(result.Data.Strategies))
	for _, s := range result.Data.Strategies {
		strategies = append(strategies, output.StrategyData{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	data := &output.StrategiesListData{
		Strategies: strategies,
	}

	formatter.FormatStrategies(data)
	return nil
}
