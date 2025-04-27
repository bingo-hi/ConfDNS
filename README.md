# 可配置的DNS

## 功能特色
 - 自由配置本地DNS，使用yaml语法，配置简单。
 - 支持DoH和DoT加密, 防域名解析被运营商篡改而导致无法正确访问到网站。
 - 使用go语言开发， 一包代码生成Windows下和Linux下的版本。
 - 可安装烧写到`openwrt`等可定制路由器中。
 - 提供加密DNS服务端URL，完美绕过运营商DNS。
 - 支持配置私有自签名证书
 - 强大的日志功能, 支持设置日志等级，日志保存时间，日志字时打包， 日志过期删除等

## 配置说明
> 参考模板 `confing/config.yaml` 
 - listen: 一般为127.0.0.1:53, 监听本地53端口。
 - upstream_dns: 指定运营商dns服务地址(非加密), 比如8.8.8.8:53 或 114.114.114.114:53
 - block_list: 禁止解析的的网址，比如不想解析一些广告网站
 - use_doh: 开启 doh 解析。当设置为true时， 默认所有请求都会通过https加密解析。当为false或没有设置时通过其它方式进行解析。

## 编译
推荐使用`build`目录下的脚本一键编译，开发者也可以使用命令行分步编译的方式 。

### 一、一键编译
#### windows开发环境
```bash
cd build
### 一键生成windows和linux下的服务
./build.bat all
```

#### Linux开发环境
```bash
cd build
### 一键生成windows和linux下的服务
./build.bat all
```

### 二、分步骤编译
#### 下载项目依赖
```bash
go mod tidy
```

#### 生成windows下可执行文件: 
``` bash
go env -w GOOS=windows
go build -o dnsclient.exe ./cmd/dnsclient
```

### 生成linux下可执行文件：
```bash
go env -w GOOS=linux
go build -o dnsclient ./cmd/dnsclient
```

## 运行
### windows下服务

### Linux下服务

## 联系作者
邮箱：mjsornp@gmail.com