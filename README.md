# RedCmd
常用的渗透测试命令生成 revers shell 内网渗透下载命令等
因为常用的命令太多了，记也懒得记，虽然有网页版的但是也懒得打开浏览器，不如命令行来的优雅，
就自己手敲了一个 代码非常简单，后续会增加命令。

<img width="910" alt="image" src="https://github.com/sspsec/RedCmd/assets/142762749/0b037ee5-87b2-4ddc-8f7d-0422723521e1">

## 使用方法
``` bash
go run redcmd.go -h
```

生成下载命令
``` bash
go run redcmd.go download 10.10.10.10 8080 exploit.exe shell.exe
```

生成反弹shell命令
``` bash
go run redcmd.go revshell 10.10.10.10 4444
```

生成meterperter shell的命令
``` bash
go run redcmd.go msfshell 10.10.10.10 4444
```

使用go编译后加入环境变量使用更佳
``` bash
go build redcmd.go
```
