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
	// without setting reader in size .
	s := bufio.NewReader(os.Stdin)
	// setting reader in size .
	//s := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(s)), 10, 64)
	checkError(sTemp, err)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(number int64, err error) {
	if err != nil {
		fmt.Println("Bad String")
	} else {
		fmt.Println(number)
	}
}

// example output
// go run Day16/main.go
// 45678
// 45678

// example output
// go run Day16/main.go
// fgyuu3
// Bad String
