# 用Go实现selpg时遇到的问题

## 错误检查
	- 把错误检查放在一个函数里，减少冗余
	````
	func check(e error) {
		if e != nil {

		panic(e)

		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)

		}
	}
	````

## 无输出，关闭文件

	一开时把文件读取作为一个函数，在函数中，defer file.Close();，导致读取的流都是空的，无输出

	解决方法：不要写成一个函数



## GO，文件读写

	````
	//相关包
	import (
		"bufio"
		"fmt"
		"io"
		"os"
	)
	````
- 获取文件指针
	````
	file, err := os.Open("file.go") // For read access.
	check(err)
	defer file.Close()
	````
- 文件输出
	````
	Reader := bufio.NewReader(file)
	for {
		inputString, readerError := Reader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			break
		}      
	}
	````
- 写文件

	````
		File, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
		check(err)
		defer outputFile.Close()

		outputWriter := bufio.NewWriter(File)
		outputString := "hello world!\n"

		outputWriter.WriteString(outputString)

		outputWriter.Flush()//清空缓冲区？？？？
	````


## flag、pflag

	falg为官方包
	
	pflag为falg扩展包
	
- 将flag绑定到一个变量
	````
	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}
	````
- pflag
	````
	var flagvar int

	func init() {
		flag.IntVar(&flagvar, "flagname", "f"， 1234, "help message for flagname")
	}
	````
- 命令中，其余参数
	````
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
		for i := 0; i != flag.NArg(); i++ {
			fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
		}
	````

## go程
- 创建子进程，将输出放入信道；在writeOutput中选择不同的输出对象
    ````
    go selectPageIntoPipe(outputPage) //输入
    writeOutput(outputPage)           //输出
    ````


## 新建cmd进程
	````
	lpcmd := exec.Command("lp", "-d", destination)

	//指定子程序输出为标准输出
	lpcmd.Stdout = os.Stdout
	lpcmd.Stderr = os.Stderr

	lpPipe, err := lpcmd.StdinPipe()
	check(err)
	//子程序开始运行
	lpcmd.Start()
	````
	
## (其他内容)[https://github.com/lianghw001/hello_go/blob/master/%E5%AD%A6%E4%B9%A0GO%E9%81%87%E5%88%B0%E7%9A%84%E9%97%AE%E9%A2%98.md]
