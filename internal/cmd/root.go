package cmd

import (
	"fmt"
	"os"

	"github.com/lzwang/lskypro-cli/internal/api"
	"github.com/lzwang/lskypro-cli/internal/config"
	"github.com/lzwang/lskypro-cli/internal/output"
	"github.com/urfave/cli/v2"
)

var (
	cfgProvider *config.Provider
	apiClient   *api.Client
	formatter   output.Formatter
)

// Execute 执行 CLI
func Execute() error {
	app := &cli.App{
		Name:     "lc",
		Usage:    "Lsky Pro CLI - 图床管理工具",
		Version:  "1.0.0",
		Commands: []*cli.Command{
			loginCommand,
			logoutCommand,
			profileCommand,
			uploadCommand,
			imagesCommand,
			albumsCommand,
			strategiesCommand,
			configCommand,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Aliases: []string{"u"},
				EnvVars: []string{"LSKY_URL"},
				Usage:   "Lsky Pro 服务器地址",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "table",
				Usage:   "输出格式: table/json/plain",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "",
				Usage:   "配置文件路径",
			},
		},
		Before: beforeAction,
	}

	return app.Run(os.Args)
}

func beforeAction(cCtx *cli.Context) error {
	// 加载配置
	var err error
	cfgProvider, err = config.NewProvider(cCtx.String("config"))
	if err != nil {
		return fmt.Errorf("加载配置失败: %w", err)
	}

	cfg := cfgProvider.Get()

	// 确定服务器地址 (优先级: 命令行 > 配置文件)
	serverURL := cCtx.String("url")
	if serverURL == "" {
		serverURL = cfg.ServerURL
	}

	// 确定输出格式 (优先级: 命令行 > 配置文件)
	outputFormat := cCtx.String("output")
	if outputFormat == "table" && cfg.Output != "" {
		outputFormat = cfg.Output
	}

	// 创建 API 客户端
	apiClient = api.NewClient(serverURL, cfg.Token)

	// 创建输出格式化器
	formatter = output.NewFormatter(outputFormat, os.Stdout)

	return nil
}
