# Lsky Pro API 文档

> 提示: 当前接口文档提取自开源版 LskyPro 的网页

## 接口说明

### 接口 URL

```
https://your-lsky-server.com/api/v1
```

---

## 验证方式

当前版本接口采用 **HTTP 基本验证** 获取 token。
获取 token 后，通过设置请求 Header 标头进行验证（Bearer Token）。

示例：

```http
Authorization: Bearer 1|1bJbwlqBfnggmOMEZqXT5XusaIwqiZjCDs7r1Ob5
```

若未设置 `Authorization` 请求上传接口，则视为游客上传。

---

## 公共请求 Headers

| 字段           | 类型     | 说明                              |
| -------------- | -------- | --------------------------------- |
| `Authorization` | String   | 授权 Token，例如：`Bearer 1`      |
| `*Accept`      | String   | 必须设置为 `application/json`     |

---

## 公共响应 Headers

| 字段                    | 类型    | 说明                     |
| ----------------------- | ------- | ------------------------ |
| `X-RateLimit-Limit`     | Integer | 当前客户端一分钟内请求配额 |
| `X-RateLimit-Remaining` | Integer | 当前客户端剩余请求配额    |

---

## HTTP 状态码说明

| 状态码 | 说明                 |
| ------ | -------------------- |
| 401    | 未登录或授权失败     |
| 403    | 管理员关闭接口功能   |
| 429    | 超出请求配额         |
| 500    | 服务端异常           |

> 文档中使用 `*` 标记的字段为必填项。

---

# 授权相关

## 生成 Token

### `POST /tokens`

### 请求参数（Body）

| 字段        | 类型   | 说明   |
| ----------- | ------ | ------ |
| `*email`    | String | 邮箱   |
| `*password` | String | 密码   |

### 返回参数

| 字段      | 类型     | 说明     |
| --------- | -------- | -------- |
| `status`  | Boolean  | 状态     |
| `message` | String   | 描述信息 |
| `data`    | Object   | 数据     |
| `token`   | String   | Token    |

---

## 清空 Token

### `DELETE /tokens`

### 返回参数

| 字段      | 类型     | 说明     |
| --------- | -------- | -------- |
| `status`  | Boolean  | 状态     |
| `message` | String   | 描述信息 |
| `data`    | Object   | 数据     |

---

# 用户资料

## `GET /profile`

### 返回参数

| 字段             | 类型     | 说明           |
| ---------------- | -------- | -------------- |
| `status`         | Boolean  | 状态           |
| `message`        | String   | 描述信息       |
| `data`           | Object   | 数据           |
| `name`           | String   | 用户名         |
| `avatar`         | String   | 头像地址       |
| `email`          | String   | 邮箱地址       |
| `capacity`       | Float    | 总容量         |
| `used_capacity`  | Float    | 已使用容量     |
| `url`            | String   | 个人主页地址   |
| `image_num`      | Integer  | 图片数量       |
| `album_num`      | Integer  | 相册数量       |
| `registered_ip`  | String   | 注册 IP        |

---

# 策略相关

## 策略列表

### `GET /strategies`

### 请求参数（Query）

| 字段      | 类型   | 说明       |
| --------- | ------ | ---------- |
| `keyword` | String | 筛选关键字 |

### 返回参数

| 字段         | 类型       | 说明     |
| ------------ | ---------- | -------- |
| `status`     | Boolean    | 状态     |
| `message`    | String     | 描述信息 |
| `data`       | Object     | 数据     |
| `strategies` | Object[]   | 策略数据 |
| `id`         | Integer    | 策略 ID  |
| `name`       | String     | 策略名称 |

---

# 图片相关

## 上传图片

### `POST /upload`

### Headers

| 字段            | 类型   | 说明                           |
| --------------- | ------ | ------------------------------ |
| `*Content-Type` | String | 必须设置为 `multipart/form-data` |

### 请求参数（Body）

| 字段          | 类型    | 说明          |
| ------------- | ------- | ------------- |
| `*file`       | File    | 图片文件      |
| `strategy_id` | Integer | 储存策略 ID   |

### 返回参数

| 字段                   | 类型     | 说明             |
| ---------------------- | -------- | ---------------- |
| `status`               | Boolean  | 状态             |
| `message`              | String   | 描述信息         |
| `data`                 | Object   | 数据             |
| `key`                  | String   | 图片唯一密钥     |
| `name`                 | String   | 图片名称         |
| `pathname`             | String   | 图片路径名       |
| `origin_name`          | String   | 原始文件名       |
| `size`                 | Float    | 图片大小（KB）   |
| `mimetype`             | String   | 图片类型         |
| `extension`            | String   | 扩展名           |
| `md5`                  | String   | MD5 值           |
| `sha1`                 | String   | SHA1 值          |
| `links`                | Object   | 链接             |
| `url`                  | String   | 访问 URL         |
| `html`                 | String   | HTML 链接        |
| `bbcode`               | String   | BBCode           |
| `markdown`             | String   | Markdown         |
| `markdown_with_link`   | String   | Markdown（带链接）|
| `thumbnail_url`        | String   | 缩略图 URL       |

---

## 图片列表

### `GET /images`

### 请求参数（Query）

| 字段         | 类型    | 说明                                  |
| ------------ | ------- | ------------------------------------- |
| `page`       | Integer | 页码                                  |
| `order`      | String  | `newest` / `earliest` / `utmost` / `least` |
| `permission` | String  | `public` / `private`                  |
| `album_id`   | Integer | 相册 ID                               |
| `keyword`    | String  | 筛选关键字                            |

### 返回参数

| 字段           | 类型       | 说明                |
| -------------- | ---------- | ------------------- |
| `status`       | Boolean    | 状态                |
| `message`      | String     | 描述信息            |
| `data`         | Object     | 数据                |
| `current_page` | Integer    | 当前页码            |
| `last_page`    | Integer    | 最后一页            |
| `per_page`     | Integer    | 每页数量            |
| `total`        | Integer    | 总数量              |
| `data`         | Object[]   | 图片列表            |
| `key`          | String     | 图片唯一密钥        |
| `name`         | String     | 图片名称            |
| `origin_name`  | String     | 原始名称            |
| `pathname`     | String     | 路径名              |
| `size`         | Float      | 大小（KB）          |
| `width`        | Integer    | 宽度                |
| `height`       | Integer    | 高度                |
| `md5`          | String     | MD5                 |
| `sha1`         | String     | SHA1                |
| `human_date`   | String     | 友好时间            |
| `date`         | String     | `yyyy-MM-dd HH:mm:ss` |
| `links`        | Object     | 链接                |

---

## 删除图片

### `DELETE /images/:key`

### 请求参数（Params）

| 字段   | 类型   | 说明     |
| ------ | ------ | -------- |
| `*key` | String | 图片密钥 |

### 返回参数

| 字段      | 类型     | 说明     |
| --------- | -------- | -------- |
| `status`  | Boolean  | 状态     |
| `message` | String   | 描述信息 |
| `data`    | Object   | 数据     |

---

# 相册相关

## 相册列表

### `GET /albums`

### 请求参数（Query）

| 字段      | 类型    | 说明                                      |
| --------- | ------- | ----------------------------------------- |
| `page`    | Integer | 页码                                      |
| `order`   | String  | `newest` / `earliest` / `most` / `least`  |
| `keyword` | String  | 筛选关键字                                |

### 返回参数

| 字段           | 类型       | 说明       |
| -------------- | ---------- | ---------- |
| `status`       | Boolean    | 状态       |
| `message`      | String     | 描述信息   |
| `data`         | Object     | 数据       |
| `current_page` | Integer    | 当前页     |
| `last_page`    | Integer    | 最后一页   |
| `per_page`     | Integer    | 每页数量   |
| `total`        | Integer    | 总数量     |
| `data`         | Object[]   | 相册列表   |
| `id`           | Integer    | 相册 ID    |
| `name`         | String     | 相册名称   |
| `intro`        | String     | 相册简介   |
| `image_num`    | Integer    | 图片数量   |

---

## 删除相册

### `DELETE /albums/:id`

### 请求参数（Params）

| 字段  | 类型   | 说明        |
| ----- | ------ | ----------- |
| `*id` | String | 相册自增 ID |

### 返回参数

| 字段      | 类型     | 说明     |
| --------- | -------- | -------- |
| `status`  | Boolean  | 状态     |
| `message` | String   | 描述信息 |
| `data`    | Object   | 数据     |
