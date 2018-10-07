````

````

## 错误检查


````
func check(e error) {
    if e != nil {
        panic(e)
    }
}
````

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

