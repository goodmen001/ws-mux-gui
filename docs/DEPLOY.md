# 部署说明

## 一、准备
- Go >= 1.21
- Windows 10/11
- GitHub Actions

## 二、本地调试
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
wails dev
```

## 三、自动构建
```bash
git tag v1.0.0
git push origin v1.0.0
```
