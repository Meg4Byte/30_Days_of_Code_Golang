package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Please input length of Array: ")
	readerArraylength := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arraylength, err := strconv.ParseInt(strings.TrimSpace(readLine(readerArraylength)), 10, 64)

	checkError(err)
	fmt.Printf("Please input String Array: ")
	readerStringArray := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stringArray := strings.Split(readLine(readerStringArray), " ")
	var numberArray = make([]int, int(arraylength))

	for k, v := range stringArray {
		n, _ := strconv.Atoi(v)
		numberArray[k] = n
	}

	sort.Ints(numberArray)
	var maxlength int
	maxlength = numberArray[len(numberArray)-1] - numberArray[0]
	if maxlength < 0 {
		maxlength = -maxlength
	}

	fmt.Println(maxlength)

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
