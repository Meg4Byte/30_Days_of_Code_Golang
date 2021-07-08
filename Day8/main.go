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
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	scanner := bufio.NewScanner(os.Stdin)
	phoneBook := make(map[string]string)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	checkError(err)

	for i := 0; i < n; i++ {
		scanner.Scan()
		arrTemp := strings.Split(scanner.Text(), " ")
		phoneBook[arrTemp[0]] = arrTemp[1]
	}

	// test case 1 query not equal input Given  names and phone numbers.
	for scanner.Scan() {
		mapName := scanner.Text()
		if x, found := phoneBook[mapName]; found {
			fmt.Printf("%v=%v\n", mapName, x)
		} else {
			fmt.Println("Not found")
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
