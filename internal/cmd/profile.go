package cmd

import (
	"fmt"

	"github.com/lzwang/lskypro-cli/internal/output"
	"github.com/urfave/cli/v2"
)

var profileCommand = &cli.Command{
	Name:    "profile",
	Aliases: []string{"pro"},
	Usage:   "查看用户资料",
	Action:  doProfile,
}

func doProfile(cCtx *cli.Context) error {
	// 检查是否已登录
	if apiClient.GetToken() == "" {
		return fmt.Errorf("请先登录: lc login")
	}

	// 调用 API 获取用户资料
	result, err := apiClient.GetProfile()
	if err != nil {
		return fmt.Errorf("获取用户资料失败: %w", err)
	}

	if !result.Status {
		return fmt.Errorf("获取用户资料失败: %s", result.Message)
	}

	// 格式化输出
	data := &output.ProfileData{
		ID:           result.Data.ID,
		Email:        result.Data.Email,
		Name:         result.Data.Name,
		Avatar:       result.Data.Avatar,
		StorageUsed:  result.Data.StorageUsed,
		ImageNum:     result.Data.ImageNum,
		AlbumNum:     result.Data.AlbumNum,
		RegisteredIP: result.Data.RegisteredIP,
	}

	formatter.FormatProfile(data)
	return nil
}
