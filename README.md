# dandanplay api wrapper
用于向dandanplay开放平台发送请求时附加身份验证信息，避免密钥泄露/公开

## 编译
1. 下载安装Go: https://go.dev/doc/install
2. 下载源码
3. 在源码所在目录下运行命令:`go build -ldflags="-s -w" -o dandanplay main.go`，其中`-o dandanplay`指编译为名称为dandanplay的可执行文件，`-ldflags="-s -w"`用于缩减可执行文件体积。
