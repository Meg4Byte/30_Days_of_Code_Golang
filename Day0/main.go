package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	fmt.Println("Hello, World.")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
