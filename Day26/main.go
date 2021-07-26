package main

import (
	"fmt"
)

func main() {
	returndate := make([]int64, 3)
	duedate := make([]int64, 3)
	for k := range returndate {
		fmt.Scan(&returndate[k])
	}
	for k := range duedate {
		fmt.Scan(&duedate[k])
	}

	res := fineCalculator(returndate[0], returndate[1], returndate[2], duedate[0], duedate[1], duedate[2])
	fmt.Println(res)

}

func fineCalculator(actualDay, actualMonth, actualYear, expectedDay, expectedMonth, expectedYear int64) int64 {
	if actualYear < expectedYear {
		return 0
	}
	if actualYear > expectedYear {
		return 10000
	}
	if actualMonth < expectedMonth {
		return 0
	}

	if actualMonth > expectedMonth {
		return (actualMonth - expectedMonth) * 500
	}

	if actualDay < expectedDay {
		return 0
	}

	if actualDay > expectedDay {
		return (actualDay - expectedDay) * 15
	}
	return 0
}
