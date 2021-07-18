package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Queue struct {
	queue []string
}

func (q *Queue) enqueue(str string) {
	q.queue = append(q.queue, str)
}

func (q *Queue) dequeue() string {
	v := q.queue[0]
	q.queue[0] = ""
	q.queue = q.queue[1:]
	return v
}

type Stack struct {
	stack []string
}

func (s *Stack) push(str string) {
	s.stack = append(s.stack, str)
}

func (s *Stack) pop() string {
	v := s.stack[len(s.stack)-1]
	s.stack[len(s.stack)-1] = ""
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

var isPalindrome bool = true

func main() {
	s := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	str := readLine(s)
	sArray := strings.Split(str, "")
	ptrQ := Queue{queue: []string{}}
	ptrS := Stack{stack: []string{}}
	for _, v := range sArray {
		ptrQ.enqueue(v)
		ptrS.push(v)
	}

	for i := 0; i < len(sArray); i++ {
		if ptrQ.dequeue() != ptrS.pop() {
			isPalindrome = false
			break
		}
		//
		if float64(i) == math.Round(float64(len(sArray))) {
			break
		}
	}

	if isPalindrome {
		fmt.Printf("The word, %v, is a palindrome.\n", str)
	} else {
		fmt.Printf("The word, %v, is not a palindrome.\n", str)
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
