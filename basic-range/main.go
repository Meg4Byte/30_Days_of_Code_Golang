package main

import (
	"fmt"
)

// source https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/

func main() {
	// Basic for-each loop (slice or array)
	fmt.Println("Basic for-each loop (slice or array)")
	a := []string{"string1", "string2", "string3"}
	for index, value := range a {
		fmt.Printf("%v %v\n", index, value)
	}
	// output indexs and values.
	// 0 string1
	// 1 string2
	// 2 string3
	a = append(a, "append value")
	for _, value := range a {
		fmt.Printf("%v \n", value)
	}
	// output: only value.
	// string1
	// string2
	// string3

	// String iteration: runes or bytes
	fmt.Println("String iteration: runes or bytes")
	for _, ch := range "天氣好" {
		fmt.Printf("%#U\n", ch)
	}
	// output
	// U+5929 '天'
	// U+6C23 '氣'
	// U+597D '好'

	// Map iteration: keys and values
	fmt.Println("Map iteration: keys and values")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	// output random map key
	// one 1
	// two 2
	// three 3
	for key, value := range m {
		if key == "one" {
			fmt.Println("map key equal one and print ")
			fmt.Println(key, value)
			delete(m, "one")
			fmt.Printf("%v\n", m)
		}
	}
}
