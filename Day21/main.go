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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// ParseInt for get number list.
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	var numberArray []int32
	for i := 0; int32(i) < n; i++ {
		nreader := bufio.NewReader(os.Stdin)
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(nreader)), 10, 64)
		checkError(err)
		numberArray = append(numberArray, int32(nTemp))
	}

	// ParseInt for get string list.
	sreader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(sreader)), 10, 64)
	checkError(err)
	s := int32(sTemp)
	var stringArray []string
	for i := 0; int32(i) < s; i++ {
		sreader := bufio.NewReader(os.Stdin)
		sTemp := readLine(sreader)
		checkError(err)
		stringArray = append(stringArray, sTemp)
	}

	printnumberArray(numberArray)
	printstringArray(stringArray)
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

func printstringArray(x []string) {
	for _, v := range x {
		fmt.Println(v)
	}
}

func printnumberArray(x []int32) {
	for _, v := range x {
		fmt.Println(v)
	}
}
