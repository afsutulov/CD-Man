@echo off
setlocal
pushd "%~dp0"
if errorlevel 1 exit /b 1

where go >nul 2>nul
if errorlevel 1 (
    echo ERROR: Go 1.22 or newer is required. Install Go and open a new terminal.
    goto :fail
)

set "CGO_ENABLED=0"
if not exist "build" mkdir "build"
if errorlevel 1 goto :fail

echo Running tests...
go test ./...
if errorlevel 1 goto :fail

echo Building CD-Man...
go build -buildvcs=false -trimpath -ldflags "-H=windowsgui" -o "build\CDMan.exe" ./cmd/cdman
if errorlevel 1 goto :fail

echo.
echo Build complete: build\CDMan.exe
popd
exit /b 0

:fail
echo Build failed. See the error above.
popd
exit /b 1
