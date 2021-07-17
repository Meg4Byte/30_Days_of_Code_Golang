package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Calculator struct {
	n          float64
	p          float64
	isNegative bool
}

func (c *Calculator) CalculatorResult() {
	if c.n < 0 || c.p < 0 {
		c.isNegative = true
		fmt.Println("n and p should be non-negative")
	}

	if c.isNegative == false {
		fmt.Println(math.Pow(c.n, c.p))
	}

}

func main() {
	s := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(s)), 10, 64)
	checkError(err)

	for i := 0; i < int(sTemp); i++ {
		s := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
		// get string like 1 2 => [1 2] []string
		arrRowTemp := strings.Split(strings.TrimRight(readLine(s), " \t\r\n"), " ")

		var arrRow []float64
		// let [1 2] []string => [1 2] []float64
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseFloat(arrRowItem, 10)
			checkError(err)
			arrRow = append(arrRow, arrItemTemp)
		}
		c := Calculator{n: arrRow[0], p: arrRow[1]}
		c.CalculatorResult()
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
		fmt.Println("Bad String")
	}
}

// example input and output

// go run Day17/main.go
// input:  4
// input:  1 2
// output: 1
// input: 3 3
// output: 27
// input: -1 9
// output: n and p should be non-negative
// input: -1 -300
// output: n and p should be non-negative
