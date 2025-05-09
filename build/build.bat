@echo off
REM Get the current script directory
SET SCRIPT_DIR=%~dp0

REM Define the project directory as the parent directory of SCRIPT_DIR
SET PROJECT_DIR=%SCRIPT_DIR%..
SET OUTPUT_DIR=%SCRIPT_DIR%output
SET LINUX_DIR=%OUTPUT_DIR%\linux
SET WINDOWS_DIR=%OUTPUT_DIR%\windows

REM Ensure output directories exist
mkdir "%LINUX_DIR%"
mkdir "%WINDOWS_DIR%"

REM Compile for Linux
echo Compiling for Linux...
REM Set GOOS to linux and build the Linux executable
set CGO_ENABLED=0
go env -w GOOS=linux
go build -ldflags="-s -w -extldflags '-static'" -o "%LINUX_DIR%\dnsclient" "%PROJECT_DIR%\cmd\dnsclient"

REM Copy the config and README files to the Linux output folder
xcopy /E /I "%PROJECT_DIR%\config" "%LINUX_DIR%\config"
copy /Y "%PROJECT_DIR%\README.md" "%LINUX_DIR%"
copy /Y "%PROJECT_DIR%\README.en-US.md" "%LINUX_DIR%"
copy /Y "%PROJECT_DIR%\README.zh-CN.md" "%LINUX_DIR%"

REM Copy LICENSE to the Linux output folder
copy /Y "%PROJECT_DIR%\LICENSE" "%LINUX_DIR%"

REM Compile for Windows
echo Compiling for Windows...
REM Set GOOS to windows and build the Windows executable
set CGO_ENABLED=0
go env -w GOOS=windows
go build -ldflags="-s -w -extldflags '-static'" -o "%WINDOWS_DIR%\dnsclient.exe" "%PROJECT_DIR%\cmd\dnsclient"

REM Copy the config and README files to the Windows output folder
xcopy /E /I "%PROJECT_DIR%\config" "%WINDOWS_DIR%\config"
copy /Y "%PROJECT_DIR%\README.md" "%WINDOWS_DIR%"
copy /Y "%PROJECT_DIR%\README.en-US.md" "%WINDOWS_DIR%"
copy /Y "%PROJECT_DIR%\README.zh-CN.md" "%WINDOWS_DIR%"

REM Copy LICENSE to the Windows output folder
copy /Y "%PROJECT_DIR%\LICENSE" "%WINDOWS_DIR%"

echo Compilation complete. Linux and Windows builds are available in the 'output' directory.
pause
