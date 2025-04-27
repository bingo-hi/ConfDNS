<!-- README.zh-CN.md -->
请选择你需要的语言版本：

- 🇨🇳 [中文版](README.zh-CN.md)
- 🇺🇸 [English](README.en-US.md)

# 可配置的DNS
下载使用：https://github.com/bingo-hi/ConfDNS/releases

## 功能特色
 - 自由配置本地DNS，使用yaml语法，配置简单。
 - 支持DoH和DoT加密, 防域名解析被运营商篡改而导致无法正确访问到网站。
 - 使用go语言开发， 一包代码生成Windows下和Linux下的版本。
 - 可安装烧写到`openwrt`等可定制路由器中。
 - 提供加密DNS服务端URL，完美绕过运营商DNS。
 - 支持配置私有自签名证书
 - 强大的日志功能, 支持设置日志等级，日志保存时间，日志定时打包， 日志过期删除等

## 配置说明
> 参考模板 `confing/config.yaml` 
 - listen: 一般为127.0.0.1:53, 监听本地53端口。
 - upstream_dns: 指定运营商dns服务地址(非加密), 比如8.8.8.8:53 或 114.114.114.114:53
 - block_list: 禁止解析的的网址，比如不想解析一些广告网站
 - use_doh: 开启 doh 解析。当设置为true时， 默认所有请求都会通过https加密解析。当为false或没有设置时通过其它方式进行解析。

## 编译
推荐使用`build`目录下的脚本一键编译，开发者也可以使用命令行分步编译的方式 。

### 1. 一键编译

#### Windows开发环境
```bash
cd build
# build both Windows and Linux binaries
./build.bat
```

#### Linux开发环境
```bash
cd build
### 一键生成windows和linux下的服务
./build.sh
```

### 2. 分步骤编译
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
> 运行需要`管理员权限`
测试运行可以命令行的方式运行， 比如：

```bash
./dnsclient
```

生产环境下，可以配置成服务的方式运行
### Linux下配置服务的方法
> 以ubuntu20.04为例，[示例模板](config/confdns.service)在`config/confdns.service`

1. 以管理员权限，在`/etc/systemd/system`目录下编辑`confdns.service`, 使`ExecStart`和`WorkingDirectory` 指向dnsclient这个可执行文件在服务器上的存放路径
2. 设置服务为开机启动
```bash
systemctl enable confdns.service
``` 
3. 启动服务
```bash
systemctl start confdns.service
```

### windows下配置服务的方法
> 以windows10为例. 本项目使用自定义服务的方式来配置管理Windows下的DNS服务

1. 生成服务exe (也可build目录下的脚本一键生成)
```bash
go env -w GOOS=windows
go build -o ConfDnsService.exe ./cmd/winservice
```

2. 安装服务（需要命令行权限打开窗口）
```bash
ConfDnsService.exe /install
```

3. 查看服务状态
```bash
sc query ConfDNS
```

4. 启动服务
```bash
sc start ConfDNS
```

5. 将网络适配器中`首选DNS服务器`改为：`127.0.0.1`

## DNS缓存
开启ConfDNS后， 若没有立即生效， 有可能是操作系统DNS缓存的原因，清理一下dns缓存过几分钟即可生效。

# 开源许可
MIT

# 联系作者
邮箱：mjsornp@gmail.com