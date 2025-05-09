#!/bin/bash

# Get the current script directory
SCRIPT_DIR=$(dirname "$(realpath "$0")")

# Define the project directory as the parent directory of SCRIPT_DIR
PROJECT_DIR=$(dirname "$SCRIPT_DIR")

# Output directory
OUTPUT_DIR="${SCRIPT_DIR}/output"
LINUX_DIR="${OUTPUT_DIR}/linux"
WINDOWS_DIR="${OUTPUT_DIR}/windows"

# Ensure output directories exist
mkdir -p "$LINUX_DIR" "$WINDOWS_DIR"

# Compile for Linux
echo "Compiling for Linux..."
# Set GOOS to linux and build the Linux executable
go env -w GOOS=linux
CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -o "$LINUX_DIR/dnsclient" "$PROJECT_DIR/cmd/dnsclient"

# Copy the config and README files to the Linux output folder
cp -r "${PROJECT_DIR}/config" "$LINUX_DIR"
cp -f "${PROJECT_DIR}/README.md" "$LINUX_DIR"
cp -f "${PROJECT_DIR}/README.en-US.md" "$LINUX_DIR"
cp -f "${PROJECT_DIR}/README.zh-CN.md" "$LINUX_DIR"

# Copy LICENSE to the Linux output folder
cp -f "${PROJECT_DIR}/LICENSE" "$LINUX_DIR"

# Compile for Windows
echo "Compiling for Windows..."
# Set GOOS to windows and build the Windows executable
go env -w GOOS=windows
CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -o "$WINDOWS_DIR/dnsclient.exe" "$PROJECT_DIR/cmd/dnsclient"

# Copy the config and README files to the Windows output folder
cp -r "${PROJECT_DIR}/config" "$WINDOWS_DIR"
cp -f "${PROJECT_DIR}/README.md" "$WINDOWS_DIR"
cp -f "${PROJECT_DIR}/README.en-US.md" "$WINDOWS_DIR"
cp -f "${PROJECT_DIR}/README.zh-CN.md" "$WINDOWS_DIR"

# Copy LICENSE to the Windows output folder
cp -f "${PROJECT_DIR}/LICENSE" "$WINDOWS_DIR"

echo "Compilation complete. Linux and Windows builds are available in the 'output' directory."
