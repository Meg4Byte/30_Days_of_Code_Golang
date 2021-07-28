package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type task struct {
	fristName string
	mail      string
}

func checkvaildName(t task) bool {
	var isStringAlphabetic = regexp.MustCompile(`^[a-z]+$`).MatchString
	if isStringAlphabetic(t.fristName) {
		return true
	} else {
		return false
	}
}

func checkvaildNamelength(t task) bool {
	if len(t.fristName) > 20 {
		return false
	} else {
		return true
	}
}

func checkvaildEmailIdlength(t task) bool {
	if len(t.mail) > 50 {
		return false
	} else {
		return true
	}
}

func checkvaildEmailIdEnd(t task) bool {
	mail := strings.Split(t.mail, "@")

	if mail[1] == "gmail.com" {
		return true
	} else {
		return false
	}
}

func main() {

	var input int64
	fmt.Scan(&input)
	inputList := make([]task, input)
	var ansList []string

	for k := range inputList {
		fmt.Scanln(&inputList[k].fristName, &inputList[k].mail)
	}
	for _, v := range inputList {
		if checkvaildName(v) && checkvaildNamelength(v) && checkvaildEmailIdlength(v) && checkvaildEmailIdEnd(v) {
			ansList = append(ansList, v.fristName)
		}
	}
	sort.Strings(ansList)

	for _, v := range ansList {
		fmt.Println(v)
	}
}
