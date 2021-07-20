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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}
	var numSwaps int
	for i := 0; int32(i) < n-1; i++ {
		for j := 0; j < int(n)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = swap(a[j], a[j+1])
				numSwaps++
			}
		}
		if numSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numSwaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[n-1])
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

func swap(x, y int32) (int32, int32) {
	return y, x
}
