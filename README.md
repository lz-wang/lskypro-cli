# lskypro-cli

Lsky Pro 图床命令行管理工具，支持图片上传、管理等功能。

## 安装

### 下载预编译的二进制文件

请访问 [Release](https://github.com/lz-wang/lskypro-cli/releases/latest) 页面下载对应的平台版本即可

### 从源码编译

```bash
git clone https://github.com/lzwang/lskypro-cli.git
cd lskypro-cli
make build
```

## 快速开始

### 1. 配置服务器地址

```bash
# 方式一：通过环境变量
export LSKY_URL="https://your-lsky-server.com/api/v1"

# 方式二：通过命令行参数
lc -u https://your-lsky-server.com/api/v1

# 方式三：通过配置文件（推荐）
lc config set url https://your-lsky-server.com/api/v1
```

### 2. 登录

```bash
lc login
# 按提示输入邮箱和密码
```

### 3. 上传图片

```bash
lc upload image.png
```

## 命令列表

### 全局选项

| 选项 | 简写 | 说明 |
|------|------|------|
| `--url` | `-u` | Lsky Pro 服务器地址 |
| `--output` | `-o` | 输出格式: table/json/plain |
| `--config` | `-c` | 配置文件路径 |

### 登录

```bash
lc login
```

登录到 Lsky Pro 服务器，按提示输入邮箱和密码。

### 登出

```bash
lc logout
```

清除本地保存的 Token。

### 用户资料

```bash
lc profile
```

查看当前登录用户的资料信息。

### 上传图片

```bash
lc upload <文件路径> [选项]
```

| 选项 | 简写 | 说明 |
|------|------|------|
| `--strategy` | `-s` | 存储策略 ID |
| `--copy` | | 复制链接格式: url/markdown/bbcode/html (默认: url) |

示例：

```bash
# 上传图片（默认复制 URL 到剪贴板）
lc upload photo.png

# 上传并复制 Markdown 格式链接
lc upload --copy markdown photo.png

# 指定存储策略
lc upload -s 2 photo.png
```

### 图片管理

```bash
# 列出图片
lc images list [选项]

# 别名
lc img ls
```

| 选项 | 简写 | 说明 |
|------|------|------|
| `--page` | `-p` | 页码 (默认: 1) |
| `--order` | | 排序: newest/earliest/utmost/least |
| `--permission` | | 权限: public/private |
| `--album` | `-a` | 相册 ID |
| `--keyword` | `-k` | 搜索关键词 |

```bash
# 删除图片
lc images delete <key> [key...] [选项]
```

| 选项 | 简写 | 说明 |
|------|------|------|
| `--force` | `-f` | 跳过确认 |
| `--yes` | `-y` | 跳过确认 |

### 相册管理

```bash
# 列出相册
lc albums list

# 别名
lc alb ls

# 删除相册
lc albums delete <id>
```

### 存储策略

```bash
# 列出存储策略
lc strategies list

# 别名
lc stg ls
```

### 配置管理

```bash
# 查看当前配置
lc config show

# 设置配置项
lc config set <key> <value>

# 可配置项:
#   url    - 服务器地址
#   output - 默认输出格式
```

## 配置文件

配置文件默认位于 `~/.lc/config.yaml`，包含以下内容：

```yaml
url: https://your-lsky-server.com/api/v1
token: your-token-here
output: table
```

## 输出格式

支持三种输出格式：

- `table` - 表格格式（默认）
- `json` - JSON 格式
- `plain` - 纯文本格式

示例：

```bash
lc -o json profile
lc -o plain images list
```

## 许可证

MIT License
