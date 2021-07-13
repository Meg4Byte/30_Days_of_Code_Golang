package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type myBook struct {
	title  string
	author string
	price  int
}

func (p *myBook) display() {
	fmt.Printf("Title: %v\n", p.title)
	fmt.Printf("Author: %v\n", p.author)
	fmt.Printf("Price: %v\n", p.price)
}

func main() {
	readerTitle := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	title := readLine(readerTitle)
	readerAuthor := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	author := readLine(readerAuthor)
	readerPrice := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	idTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(readerPrice)), 10, 64)
	checkError(err)
	price := int(idTemp)
	p := myBook{title, author, price}
	p.display()

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
