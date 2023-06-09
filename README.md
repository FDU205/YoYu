# ![幽语YOYU App](logo.png)

# [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/FDU205/YoYu/blob/main/LICENSE)
[![GoDoc](https://godoc.org/github.com/gothinkster/golang-gin-realworld-example-app?status.svg)](https://godoc.org/github.com/gothinkster/golang-gin-realworld-example-app)

# 后端快速启动（仅介绍Linux下的使用，其他系统下请自行探索）

## 1、安装Redis

首先安装[Redis](https://redis.io/)数据库, 点击链接去官网下载或使用系统自带的包管理工具安装

```
ubuntu: apt install redis
arch linux: pacman -S redis
...
```

安装完毕后使用`redis-server`命令运行，如有输出则说明安装正确

## 2、安装MySQL

点击进入[MySQL](https://www.mysql.com/)官网下载，或使用系统自带的包管理工具安装

```
ubuntu: apt install mysql
arch linux自带无需安装
...
```

使用`mysql -uroot `命令进入 mysql，能正常进入则说明安装正确

在root用户下使用如下命令配置数据库

```
create database yoyu;
grant all on yoyu.* to yoyu@'localhost' identified by '123456';
exit;
```

接着输入`mysql -uyoyu -p123456`命令，能正常进入则说明配置成功。

## 3、安装Golang

如果你的机器上没有Golang或版本低于 Go 1.13，请前往Go官网安装最新版

https://golang.org/doc/install

## 4、从源码运行

首先启动Redis数据库

在[main.go](./backend/main.go)的最后一行修改你需要的ip和port

到目录[backend](./backend)下使用 `go run main.go`命令运行

或 使用 `go build`编译出二进制文件然后进入[release](./backend/release)执行

tips: 在backend目录下使用`go test`命令可以执行所有的单元测试

## 更多

更多内容请查看

[1、后端设计文档](./docs/后端设计文档.md)

[2、前端设计文档](./docs/前端设计文档.md)

[3、API设计文档](./docs/API设计文档.md)