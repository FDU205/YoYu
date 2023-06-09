# 幽语YOYU<img src="./logo.png" width="5%">


# [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/FDU205/YoYu/blob/main/LICENSE) 

# 快速启动（仅介绍Linux下的使用，其他系统下请自行探索）
## 前端快速启动

### 1、安装Node.js

首先安装[Node.js](https://nodejs.org/en/)，点击链接去官网下载或使用系统自带的包管理工具安装

```bash
ubuntu: apt install nodejs
arch linux: pacman -S nodejs
...
```

安装完毕后使用`node -v`命令查看版本，如有输出则说明安装正确

### 2、安装yarn

使用`npm install -g yarn`命令安装yarn

安装完毕后使用`yarn -v`命令查看版本，如有输出则说明安装正确

### 3、安装依赖

进入目录[my_app](./frontend/my_app/)下

使用`yarn install`命令安装依赖

### 4、运行

使用`expo start`命令运行 <br>
手机上安装[Expo](https://expo.io/)，扫描二维码即可运行

## 后端快速启动

### 1、安装Redis

首先安装[Redis](https://redis.io/)数据库, 点击链接去官网下载或使用系统自带的包管理工具安装

```
ubuntu: apt install redis
arch linux: pacman -S redis
...
```

安装完毕后使用`redis-server`命令运行，如有输出则说明安装正确

### 2、安装MySQL

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

### 3、安装Golang

如果你的机器上没有Golang或版本低于 Go 1.13，请前往Go官网安装最新版

https://golang.org/doc/install

### 4、从源码运行

首先启动Redis数据库

在[main.go](./backend/main.go)的最后一行修改你需要的ip和port

进入目录[backend](./backend)下

使用 `go run main.go`命令运行

或 使用 `go build`编译出二进制文件然后使用`./backend`命令执行

## 更多

更多内容请查看

[1、后端设计文档](./docs/后端设计文档.md)

[2、前端设计文档](./docs/前端设计文档.md)

[3、API设计文档](./docs/API设计文档.md)

[4、功能设计文档](./docs/功能设计文档.md)
