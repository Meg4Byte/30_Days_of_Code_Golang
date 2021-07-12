package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type person struct {
	firstName  string
	lastName   string
	id         int
	numberList []int
}

func (p *person) PrintPerson() {
	fmt.Printf("Name: %v, %v\n", p.lastName, p.firstName)
	fmt.Printf("ID: %v\n", p.id)
}

func (p *person) PrintCalculate() {
	var finscore int
	var sumNum int
	var finLettle string
	for _, v := range p.numberList {
		sumNum += v
	}
	finscore = sumNum / len(p.numberList)
	switch {
	case finscore >= 90:
		finLettle = "O"
	case finscore >= 80:
		finLettle = "E"
	case finscore >= 70:
		finLettle = "A"
	case finscore >= 55:
		finLettle = "P"
	case finscore >= 40:
		finLettle = "T"
	default:
		fmt.Println("Not Weird")
	}

	fmt.Printf("Grade: %v\n", finLettle)
}

func main() {
	readerf := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	f := readLine(readerf)
	readerl := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	l := readLine(readerl)
	readerid := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	idTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(readerid)), 10, 64)
	checkError(err)
	id := int(idTemp)

	readerNumList := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(readLine(readerNumList)), " ")

	var realNumList = []int{}
	for _, i := range arrTemp {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		realNumList = append(realNumList, j)
	}
	p := &person{firstName: f, lastName: l, id: id, numberList: realNumList}
	p.PrintPerson()
	p.PrintCalculate()

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
