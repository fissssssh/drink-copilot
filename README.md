# drink-copilot

一个基于 GitHub Actions 的喝水提醒小工具：

- 每天北京时间 **09:00 ~ 23:00** 每小时自动发送一次提醒
- 支持手动触发发送

## 请求说明

本项目发送提醒使用如下请求格式：

`https://<uid>.push.ft07.com/send/<sendkey>.send?title=<title>&desp=<desp>`

请求参数与返回说明参考 [Server酱³ 官方文档](https://doc.sc3.ft07.com/zh/serverchan3/server/api)

## 项目结构

- `main.go`：发送提醒的 Go 程序（每运行一次发送一次）
- `.github/workflows/send-once.yml`：定时/手动触发的 GitHub Actions 工作流

## 使用前准备

在 GitHub 仓库中配置以下 Actions Secrets：

- `UID`：你的 uid
- `SENDKEY`：你的 sendkey

路径：`Settings -> Secrets and variables -> Actions`

## GitHub Actions 触发方式

工作流：`Send Water Reminder Once`

- 定时触发：`0 1-15 * * *`（UTC），对应北京时间 09:00~23:00 每小时一次
- 手动触发：在 Actions 页面点击 `Run workflow`

## 本地运行

程序依赖 Go 1.22+。

Windows PowerShell：

```powershell
$env:UID="你的uid"
$env:SENDKEY="你的sendkey"
$env:PUSH_UID=$env:UID
go run .
```

程序读取环境变量：

- `PUSH_UID`
- `SENDKEY`

当前默认文案：

- `title`: `喝水提醒`
- `desp`: `该喝水啦 💧`

## 常见问题

Q: 为什么 workflow 里关闭了 `setup-go` 缓存？

A: 项目仅使用标准库，没有 `go.sum`，关闭缓存可避免 `Restore cache failed` 报错。
