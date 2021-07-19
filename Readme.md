# Bread 启动中断工具Golang版本

![](https://img.shields.io/badge/license-MIT-000000.svg)
[![Build Status](https://www.travis-ci.org/haowanxing/breed-enter-go.svg?branch=master)](https://www.travis-ci.org/haowanxing/breed-enter-go)
![](https://img.shields.io/badge/language-Golang-green.svg)
[![](https://img.shields.io/badge/website-imsry.cn-blue.svg)](https://www.imsry.cn)

> 该工具仅仅实现了发送中断命令，以减轻手动捅"菊花"的操作负担。

关于Breed请移步：[hackpascal的Breed发布帖](https://www.right.com.cn/forum/thread-161906-1-1.html) 、[Breed简单使用以及常见问题解答](https://www.right.com.cn/forum/thread-174525-1-1.html)

## 编译

> 编译环境：go version go1.16.3 darwin/amd64

编译命令：

```bash
go build -o BreedEnter main.go
```

## 运行

已发布的编译版本：[Releases1.0](https://github.com/haowanxing/breed-enter-go/releases/tag/v1.0)

### Linux环境：

执行编译好的二进制文件：

```bash
./BreedEnter
```

### Windows环境：

在CMD或PowerShell中运行:

```bash
breed-enter_windows_amd64.exe
```