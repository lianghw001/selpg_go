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

- 测试的输入文件
[inputTest.txt](https://github.com/lianghw001/selpg_go/blob/master/inputTest.txt)

[inputTest2.txt](https://github.com/lianghw001/selpg_go/blob/master/inputTest2.txt)

- 1.
  ````
  selpg -s1 -e1 inputTest.txt
  ````
  ![01](https://github.com/lianghw001/selpg_go/blob/master/picture/01.PNG)
- 2.
  ````
  selpg -s1 -e1 < inputTest.txt
  ````
  ![02](https://github.com/lianghw001/selpg_go/blob/master/picture/02.PNG)
- 3.
  ````
  cat inputTest.txt | selpg -s1 -e1
  ````
  ![03](https://github.com/lianghw001/selpg_go/blob/master/picture/03.PNG)
- 4.
  ````
  selpg -s1 -e1 inputTest.txt > outputTest.txt
  ````
  ![04](https://github.com/lianghw001/selpg_go/blob/master/picture/04.PNG)
- 5. 为了有error，去除-e
  ````
  selpg -s1   inputTest.txt 2>error.txt
  ````
  ![05](https://github.com/lianghw001/selpg_go/blob/master/picture/05.PNG)
- 6.
  ````
  selpg -s1 -e1 inputTest.txt | cat
  ````
  ![06](https://github.com/lianghw001/selpg_go/blob/master/picture/06.PNG)

- 7. 每页3行
  ````
  selpg -s1 -e1 -l3 inputTest.txt
  ````
  ![07](https://github.com/lianghw001/selpg_go/blob/master/picture/07.PNG)
- 8. 以换页符换页
  ````
  selpg -s1 -e1 -f inputTest2.txt
  ````
  ![08](https://github.com/lianghw001/selpg_go/blob/master/picture/08.PNG)
- 9. 没有打印机，用cat测试
  ````
  selpg -s1 -e1 -l6 -dlp1 inputTest.txt
  ````
  ![09](https://github.com/lianghw001/selpg_go/blob/master/picture/09.PNG)


