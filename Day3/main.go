package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 給一個整數 n, 如果是奇數或者是 6~20 之間的偶數就說他是 Weird，反之就是 Not Weird。
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	if N%2 != 0 {
		fmt.Println("Weird")
	} else {
		switch {
		case 6 <= N && N <= 20:
			fmt.Println("Weird")
		default:
			fmt.Println("Not Weird")
		}
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
