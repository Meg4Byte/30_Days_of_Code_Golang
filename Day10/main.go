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
	// read stdin
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// parse reader to int.
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	// convert base10 to base2
	binaryString := (strconv.FormatInt(int64(nTemp), 2))
	// split base2 number to array via 0.
	arrayString := strings.Split(binaryString, "0")
	var max int

	// check the valus len in that array.
	for _, value := range arrayString {
		if l := len(value); l > max {
			max = l
		}
	}
	fmt.Println(max)

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
