<!-- README.en-US.md -->
Please select the language version you need:

- 🇨🇳 [中文版](README.zh-CN.md)
- 🇺🇸 [English](README.en-US.md)

# Configurable DNS Client

## Features
- Fully configurable local DNS using simple YAML syntax.
- Supports encrypted DNS via DoH and DoT to prevent ISP tampering.
- Written in Go; a single codebase builds both Windows and Linux binaries.
- Can be installed/flashed into custom routers like OpenWrt.
- Provides encrypted DNS server URLs to bypass ISP DNS.
- Supports private self-signed certificates.
- Powerful logging: set log level, retention period, scheduled archiving, expiration-based cleanup.

## Configuration
> See template at `config/config.yaml`

- **listen**: usually `127.0.0.1:53`, binds to local port 53.  
- **upstream_dns**: your ISP’s plaintext DNS server, e.g. `8.8.8.8:53` or `114.114.114.114:53`.  
- **block_list**: list of domains to block (e.g. ad domains).  
- **use_doh**: enable DoH resolution.  
  - `true`: all queries go over HTTPS.  
  - `false` or omitted: use standard DNS.

## Build

You can either run the one-click scripts in the `build` folder or build step by step.

### 1. One-click build

#### Windows
```bash
cd build
# build both Windows and Linux binaries
./build.bat all
```

#### Linux
```bash
cd build
### one click to generate service
./build.sh all
```

### 2. Step-by-step
#### Download dependencies
```bash
go mod tidy
```

#### Build Windows executable
``` bash
go env -w GOOS=windows
go build -o dnsclient.exe ./cmd/dnsclient
```

### Build Linux executable
```bash
go env -w GOOS=linux
go build -o dnsclient ./cmd/dnsclient
```

## Run
### On Windows

### On Linux

## Contact
Email: mjsornp@gmail.com