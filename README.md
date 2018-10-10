# 服务计算homework
用go语言实现selpg (学习[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html))



## 介绍

从输入文本抽取指定页，输出

- mandatory_opts强制选项
  开始的页数-sNum、结束的页数-eNum
  ````
  $ selpg -s10 -e20 ...
  ````

- optional_opts 可选选项
  - 以行数换页，设置每页的行数-lNum,以换页符\f换页-f

    lNum，-f两个选项是互斥的
    ````
    $ selpg -s10 -e20  ... ##以行数换页，每页72行（默认值）
    $ selpg -s10 -e20 -l66 ... ##以行数换页，每页66行
    $ selpg -s10 -e20 -f ...   ##以换页符\f换页
    ````
  - -dDestination,输出至打印程序
    ````
    $ selpg -s10 -e20 -dlp1 
    ````
- 其他

  页数从1开始计数
  
  使用了pflag包，需要安装pflag
  `` go get github.com/spf13/pflag ``
  
##  使用测试
1. fdsa
````
1
$ selpg -s1 -e1 input_file
该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。

1
$ selpg -s1 -e1 < input_file
该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。

1
$ other_command | selpg -s10 -e20
“other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 10 页到第 20 页写至 selpg 的标准输出（屏幕）。

1
$ selpg -s10 -e20 input_file >output_file
selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“output_file”。

1
$ selpg -s10 -e20 input_file 2>error_file
selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。请注意：在“2”和“>”之间不能有空格；这是 shell 语法的一部分（请参阅“man bash”或“man sh”）。

1
$ selpg -s10 -e20 input_file >output_file 2>error_file
selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至“error_file”。当“input_file”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。

1
$ selpg -s10 -e20 input_file >output_file 2>/dev/null
selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。

1
$ selpg -s10 -e20 input_file >/dev/null
selpg 将第 10 页到第 20 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。这可作为测试 selpg 的用途，此时您也许只想（对一些测试情况）检查错误消息，而不想看到正常输出。

1
$ selpg -s10 -e20 input_file | other_command
selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 10 页到第 20 页被写至该标准输入。“other_command”的示例可以是 lp，它使输出在系统缺省打印机上打印。“other_command”的示例也可以 wc，它会显示选定范围的页中包含的行数、字数和字符数。“other_command”可以是任何其它能从其标准输入读取的命令。错误消息仍在屏幕显示。

1
$ selpg -s10 -e20 input_file 2>error_file | other_command
与上面的示例 9 相似，只有一点不同：错误消息被写至“error_file”。

在以上涉及标准输出或标准错误重定向的任一示例中，用“>>”替代“>”将把输出或错误数据附加在目标文件后面，而不是覆盖目标文件（当目标文件存在时）或创建目标文件（当目标文件不存在时）。

以下所有的示例也都可以（有一个例外）结合上面显示的重定向或管道命令。我没有将这些特性添加到下面的示例，因为我认为它们在上面示例中的出现次数已经足够多了。例外情况是您不能在任何包含“-dDestination”选项的 selpg 调用中使用输出重定向或管道命令。实际上，您仍然可以对标准错误使用重定向或管道命令，但不能对标准输出使用，因为没有任何标准输出 — 正在内部使用 popen() 函数由管道将它输送至 lp 进程。

1
$ selpg -s10 -e20 -l66 input_file
该命令将页长设置为 66 行，这样 selpg 就可以把输入当作被定界为该长度的页那样处理。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

1
$ selpg -s10 -e20 -f input_file
假定页由换页符定界。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

1
$ selpg -s10 -e20 -dlp1 input_file
第 10 页到第 20 页由管道输送至命令“lp -dlp1”，该命令将使输出在打印机 lp1 上打印。

最后一个示例将演示 Linux shell 的另一特性：

1
$ selpg -s10 -e20 input_file > output_file 2>error_file &


