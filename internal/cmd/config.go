package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var configCommand = &cli.Command{
	Name:    "config",
	Aliases: []string{"cfg"},
	Usage:   "配置管理",
	Subcommands: []*cli.Command{
		{
			Name:    "view",
			Aliases: []string{"show"},
			Usage:   "查看当前配置",
			Action:  doConfigView,
		},
		{
			Name:    "set",
			Usage:   "设置配置项",
			Action:  doConfigSet,
		},
	},
}

func doConfigView(cCtx *cli.Context) error {
	cfg := cfgProvider.Get()

	fmt.Printf("配置文件: %s\n", cfgProvider.GetConfigPath())
	fmt.Println()
	fmt.Printf("服务器地址: %s\n", cfg.ServerURL)
	if cfg.Token != "" {
		fmt.Printf("Token: %s... (已设置)\n", cfg.Token[:20])
	} else {
		fmt.Println("Token: (未设置)")
	}
	fmt.Printf("输出格式: %s\n", cfg.Output)

	return nil
}

func doConfigSet(cCtx *cli.Context) error {
	if cCtx.NArg() < 2 {
		return fmt.Errorf("用法: lc config set <key> <value>")
	}

	key := cCtx.Args().Get(0)
	value := cCtx.Args().Get(1)

	switch key {
	case "url", "server_url":
		cfgProvider.SetServerURL(value)
	case "output":
		cfgProvider.SetOutput(value)
	default:
		return fmt.Errorf("未知的配置项: %s", key)
	}

	if err := cfgProvider.Save(); err != nil {
		return fmt.Errorf("保存配置失败: %w", err)
	}

	fmt.Printf("已设置 %s = %s\n", key, value)
	return nil
}
