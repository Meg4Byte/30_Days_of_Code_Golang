package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// start reader.
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// TrimSpace 返回字符串 s 的一部分，刪除所有前導和尾隨空格，如 Unicode 定義的那樣。
	t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// check str convert to int not error.
	checkError(err)

	for i := 0; i < int(t); i++ {
		str := readLine(reader)
		strarray := strings.Split(str, "")
		var tempOddArray []string
		var tempEvenArray []string
		for i, ch := range strarray {
			if i%2 == 0 {
				tempEvenArray = append(tempEvenArray, ch)
			} else {
				tempOddArray = append(tempOddArray, ch)
			}
		}
		for _, ch := range tempEvenArray {
			fmt.Printf("%v", ch)
		}
		fmt.Printf("%v", " ")
		for _, ch := range tempOddArray {
			fmt.Printf("%v", ch)
		}
		fmt.Println()
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
