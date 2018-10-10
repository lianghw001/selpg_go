package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

var start int
var end int
var pageLen int
var pageType bool //when pageType == 1,ignore pageLen.
var destination string

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func main() {
	//error 的 输出流
	initFlag()
	flag.Parse()
	checkCommand()
	outputPage := make(chan []byte)
	go selectPageIntoPipe(outputPage)
	writeOutput(outputPage) //输出

}
func checkCommand() {
	if start <= 0 || end <= 0 || start > end {
		fmt.Fprintln(os.Stderr, "error:-s,-e must be confirm")
		os.Exit(2)
	}
	if pageLen < 1 {
		fmt.Fprintln(os.Stderr, "pageLen must bigger than 0")
		os.Exit(3)
	}
	if pageType && pageLen != 72 {
		fmt.Fprintln(os.Stderr, "warning:When use '-f' option,'-l' will be ignored")
	}
	if flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, "warning:Only the first argument will work, others will be ignored")
	}

}
func initFlag() {

	flag.IntVarP(&start, "start", "s", -1, "strat page number")
	flag.IntVarP(&end, "end", "e", -1, "end page number")
	flag.IntVarP(&pageLen, "pageLen", "l", 72, "the line number of one page")
	flag.BoolVarP(&pageType, "pageType", "f", false, "type of page,change page with '\f' or not")
	flag.StringVarP(&destination, "destination", "d", "", "output destination")
}

func selectPageIntoPipe(outputPage chan []byte) {
	defer close(outputPage)

	var inputStr *bufio.Reader
	if flag.NArg() == 0 {
		inputStr = bufio.NewReader(os.Stdin)
		//stand input
	} else {
		file, err := os.Open(flag.Arg(0)) // For read access.
		defer file.Close()
		check(err)
		inputStr = bufio.NewReader(file)
		//file input
	}

	var delim byte
	var count int
	if pageType {
		delim = '\f'
		count = 1
	} else {
		delim = '\n'
		count = pageLen
	}
	currentCount := 0
	currentPage := 1
	var lineOrPage []byte
	var err error
	for {

		if currentCount == count {
			currentPage++
			currentCount = 0
		}
		currentCount++
		lineOrPage, err = inputStr.ReadSlice(delim)

		if err == io.EOF {
			if currentPage < start {
				fmt.Fprintln(os.Stderr, "total page is less than start page ")
				os.Exit(4)
			}
			if currentPage < end {
				fmt.Fprintln(os.Stderr, "total page is less than end page ")
				os.Exit(5)
			}
			outputPage <- lineOrPage //输出最后的不足一页的内容
			return
		}

		if currentPage >= start {
			if currentPage > end {
				break
			}
			outputPage <- lineOrPage
		}
	}

}
func writeOutput(outputPage chan []byte) {
	if destination == "" {
		writer := bufio.NewWriter(os.Stdout)
		for i := range outputPage {
			_, err := writer.Write(i) //记录博客
			check(err)
			writer.Flush()
		}
	} else {
		//“-dlp1”
		//输出至打印机
		lpcmd := exec.Command("lp", "-d", destination)

		//使用cat测试
		//lpcmd := exec.Command("cat")

		//指定子程序输出为标准输出
		lpcmd.Stdout = os.Stdout
		lpcmd.Stderr = os.Stderr

		lpPipe, err := lpcmd.StdinPipe()
		check(err)
		lpcmd.Start()

		for i := range outputPage {
			fmt.Fprintf(lpPipe, "%q", i)

		}
	}

}
