<!-- README.en-US.md -->
Please select the language version you need:

- 🇨🇳 [中文版](README.md)
- 🇺🇸 [English](README_en.md)

# Configurable DNS Client
download：https://github.com/bingo-hi/ConfDNS/releases

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
./build.bat
```

#### Linux
```bash
cd build
### one click to generate service
./build.sh
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
> Running requires **administrator privileges**.

For testing purposes, you can run it via the command line, for example:
```bash
./dnsclient
```

### How to configure the service in production environment

#### How to configure the service on Linux
> Example for Ubuntu 20.04, [sample template](config/confdns.service) located at `config/confdns.service`

1. With administrator privileges, edit `confdns.service` in the `/etc/systemd/system` directory, and make sure `ExecStart` and `WorkingDirectory` point to the path where the `dnsclient` executable file is stored on the server.
2. Set the service to start on boot:
```bash
systemctl enable confdns.service
```
3. Start the service:
```bash
systemctl start confdns.service
```

#### How to configure the service on Windows
> Example for Windows 10. This project uses a custom service method to configure and manage the DNS service on Windows.

1. Generate the service executable (or use the script in the `build` directory to generate it with one click):
```bash
go env -w GOOS=windows
go build -o ConfDnsService.exe ./cmd/winservice
```
2. Install the service (requires command-line privileges to open the window):
```bash
ConfDnsService.exe /install
```
3. Check the service status:
```bash
sc query ConfDNS
```
4. Start the service:
```bash
sc start ConfDNS
```
5. Change the "Preferred DNS server" in the network adapter to: `127.0.0.1`

## DNS Cache
After enabling ConfDNS, if it does not take effect immediately, it may be due to the operating system's DNS cache. Clearing the DNS cache and waiting a few minutes should resolve the issue.

# Source Licence
MIT

# Contact
Email: mjsornp@gmail.com