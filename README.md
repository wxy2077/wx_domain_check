# 微信域名检测

参考 https://www.zhihu.com/question/42101054

> 原理： 主要是用微信公众平台 长链接转成短链接 然后判断短链接重定向

微信公众平台测试账号链接：http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login

## 写这个API的目的

主要是为了学习Go语言，使用Gin框架，学习掌握了Go各种数据类型的转换，了解yaml配置文件读取，
模块初始化配置，net/http请求模块的封装。

目录结构

```   
|____config     
| |____config.go        // 读取配置文件
|____config.yaml        // 配置文件
|____utils              // 工具文件
| |____request.go       // 发送请求
| |____tools.go         // 工具类
|____routes
| |____check.go         // api响应函数
| |____router.go        // 区分api路由
|____main.go            // 启动入口文件
|____go.mod 
|____go.sum
|____README.md

```

## 如何运行

golang 版本最好1.13+

### 安装依赖
```shell
go mod download
```

### 运行
```
go run main.go 
```

## 线上API

测试 GET请求 

```
https://www.charmcode.cn/tools/domain/check?url=http://xxx.xxx.com/
```
正常返回 code 200

```json
{
    "code": 200,
    "data": {
        "isFreeze": true,
        "url": "http://xxx.xxx.com/"
    },
    "message": "success"
}
```

## 部署

```
// 编译
go build -o "domainAPI"
// 添加参数指定运行平台 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o "domainAPI"
```

将程序上传到Linux服务器指定的目录，执行下列命令

```
chmod 755 domainAPI     // 修改程序权限 https://stackoverflow.com/questions/18932536/ftp-rights-755-vs-777

// 将程序放在后台运行，并输出日志到指定文件(可以使用supervisor托管) 
nohup ./domainAPI >/home/log/ginapi.log 2>&1 &
```
