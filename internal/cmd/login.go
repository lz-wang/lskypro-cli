package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

var loginCommand = &cli.Command{
	Name:  "login",
	Usage: "登录获取 Token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "email",
			Aliases: []string{"e"},
			EnvVars: []string{"LSKY_EMAIL"},
			Usage:   "邮箱",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			EnvVars: []string{"LSKY_PASSWORD"},
			Usage:   "密码",
		},
		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "服务器地址 (首次使用必填)",
		},
	},
	Action: doLogin,
}

func doLogin(cCtx *cli.Context) error {
	email := cCtx.String("email")
	password := cCtx.String("password")
	url := cCtx.String("url")

	// 如果没有提供服务器地址，使用配置中的
	if url == "" {
		url = apiClient.GetBaseURL()
	}

	// 检查服务器地址
	if url == "" {
		return fmt.Errorf("请提供服务器地址: lc login -u https://your-lsky-server.com")
	}

	// 交互式输入邮箱
	if email == "" {
		fmt.Print("请输入邮箱: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("读取邮箱失败: %w", err)
		}
		email = strings.TrimSpace(input)
	}

	// 交互式输入密码
	if password == "" {
		fmt.Print("请输入密码: ")
		bytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println() // 换行
		if err != nil {
			return fmt.Errorf("读取密码失败: %w", err)
		}
		password = string(bytes)
	}

	// 设置服务器地址
	apiClient.SetBaseURL(url)

	// 调用登录 API
	result, err := apiClient.Login(email, password)
	if err != nil {
		return fmt.Errorf("登录失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("登录失败: %s", result.Message)
	}

	// 保存配置
	cfgProvider.SetServerURL(url)
	cfgProvider.SetToken(result.Data.Token)
	if err := cfgProvider.Save(); err != nil {
		return fmt.Errorf("保存配置失败: %w", err)
	}

	fmt.Printf("登录成功! Token 已保存到 %s\n", cfgProvider.GetConfigPath())
	return nil
}
