package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var logoutCommand = &cli.Command{
	Name:  "logout",
	Usage: "退出登录 (清空 Token)",
	Action: doLogout,
}

func doLogout(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		fmt.Println("当前未登录")
		return nil
	}

	// 调用 API 清空服务端 Token
	if err := apiClient.Logout(); err != nil {
		// 即使 API 调用失败，也清除本地 Token
		fmt.Printf("警告: 清空服务端 Token 失败: %v\n", err)
	}

	// 清除本地配置中的 Token
	cfgProvider.ClearToken()
	if err := cfgProvider.Save(); err != nil {
		return fmt.Errorf("保存配置失败: %w", err)
	}

	fmt.Println("已退出登录")
	return nil
}
