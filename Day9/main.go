package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'factorial' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 *
 * if input factorial(3) = 3*2*1 = 6
 * if input factorial(5) = 5*4*3*2*1 = 120
 */

// soultion 1
func factorial(n int32) int32 {
	// Write your code here
	var fin int32 = 1
	if 2 <= n && n < 12 {
		for i := n; i >= int32(1); i-- {
			fin = fin * i
		}
		return fin
	}
	fin = n
	return fin
}

// soultion 2
func factorial2(n int32) int32 {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial2(n-1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := factorial(n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
