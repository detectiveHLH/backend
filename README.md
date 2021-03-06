<h1 align="center">go-backend-starter</h1>

中文 | [English](./README-en.md)

## 介绍

使用Go Module作为依赖管理，基于Gin搭建Go的Web服务器，使用Endless来使服务器平滑重启，使用Swagger来自动生成Api文档。

## 安装

```bash
git clone https://github.com/detectiveHLH/go-backend-starter.git
cd go-backend-starter
go get
go run main.go
```

## 注意

- go的版本需要高于1.11
- 使用goland时，需要确保Go module是Enable状态

## 使用

- 查看[Swagger API文档](http://localhost:8080/swagger/index.html)
- 访问[登录](http://localhost:8080/login?username=test&password=123)接口，可以获取Access Token

## 许可

[MIT](./LICENSE)

