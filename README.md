````
页数从1开始计数
````
# selpg_go

从输入文本抽取指定页，输出

mandatory_opts强制选项
开始的页数-sNum、结束的页数-eNum
````
$ selpg -s10 -e20 ...
````

optional_opts 可选选项
以行数换页，设置每页的行数-lNum,以换页符\f换页-f
-lNum，-f两个选项是互斥的
````
$ selpg -s10 -e20  ... ##以行数换页，每页72行（默认值）
$ selpg -s10 -e20 -l66 ... ##以行数换页，每页66行
$ selpg -s10 -e20 -l ... ##以行数换页，每页72行（默认值）
$ selpg -s10 -e20 -f ...   ##以换页符\f换页
````

使用了pflag包
需要安装pflag
go get github.com/spf13/pflag
# 没看明白

“-dDestination”可选选项：
selpg 还允许用户使用“-dDestination”选项将选定的页直接发送至打印机。这里，“Destination”应该是 lp 命令“-d”选项（请参阅“man lp”）可接受的打印目的地名称。该目的地应该存在 ― selpg 不检查这一点。在运行了带“-d”选项的 selpg 命令后，若要验证该选项是否已生效，请运行命令“lpstat -t”。该命令应该显示添加到“Destination”打印队列的一项打印作业。如果当前有打印机连接至该目的地并且是启用的，则打印机应打印该输出。这一特性是用 popen() 系统调用实现的，该系统调用允许一个进程打开到另一个进程的管道，将管道用于输出或输入。在下面的示例中，我们打开到命令

1
$ lp -dDestination
的管道以便输出，并写至该管道而不是标准输出：

1
selpg -s10 -e20 -dlp1
该命令将选定的页作为打印作业发送至 lp1 打印目的地。您应该可以看到类似“request id is lp1-6”的消息。该消息来自 lp 命令；它显示打印作业标识。如果在运行 selpg 命令之后立即运行命令 lpstat -t | grep lp1 ，您应该看见 lp1 队列中的作业。如果在运行 lpstat 命令前耽搁了一些时间，那么您可能看不到该作业，因为它一旦被打印就从队列中消失了。
