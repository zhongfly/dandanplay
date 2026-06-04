# dandanplay API wrapper

用于向 dandanplay 开放平台发起请求时自动附加认证信息，避免在调用脚本中直接暴露签名细节。

## 功能
- 当请求地址以 `https://api.dandanplay.net/` 开头时，自动添加：
  - `X-AppId`
  - `X-Timestamp`
  - `X-Signature`
- 当请求地址为 `https://api.dandanplay.net/api/v2/login` 时，自动根据 `userName`、`password` 计算并补充登录参数：
  - `hash`
  - `unixTimestamp`
  - `appId`
- 支持常见的 curl 风格参数：
  - `-X` 指定 HTTP 方法
  - `-d` 指定请求体
  - `-H` 多次传入请求头
  - `-o` 输出到文件

## 编译
1. 安装 Go（建议 1.23+）：https://go.dev/doc/install
2. 下载本仓库源码。
3. 在仓库目录执行：

```bash
GO111MODULE=off go build -ldflags="-s -w" -o dandanplay main.go
```

参数说明：
- `-o dandanplay`：输出可执行文件名为 `dandanplay`
- `-ldflags="-s -w"`：减小可执行文件体积

## 配置
`main.go` 中包含以下常量：

```go
const (
    AppId     = "your_app_id"
    AppSecret = "your_app_secret"
)
```

请在编译前替换为你自己的开放平台凭据，或在 CI 中通过替换注入（本仓库工作流即采用该方式）。

> 注意：不要将真实 `AppId`/`AppSecret` 提交到公开仓库。

## 使用
命令格式：

```bash
./dandanplay [options] <URL>
```

可用参数：
- `-X`：HTTP 方法，默认 `GET`
- `-d`：请求数据（通常为 JSON 字符串）
- `-H`：请求头，可重复使用，例如 `-H "Content-Type: application/json"`
- `-o`：将响应体写入文件

### 示例 1：普通 GET 请求

```bash
./dandanplay "https://api.dandanplay.net/api/v2/bangumi/123"
```

### 示例 2：登录请求（自动补全 hash）

```bash
./dandanplay \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"userName":"your_name","password":"your_password"}' \
  "https://api.dandanplay.net/api/v2/login"
```

## 工作流说明
仓库内 GitHub Actions 会进行多平台编译（Linux/Windows/macOS）并尝试使用 UPX 压缩产物。
