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

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// count total cost, note: float can not compute with int, so need to convert type.
	var total_cost float64 = meal_cost + (meal_cost * float64(tip_percent) / 100) + (meal_cost * float64(tax_percent) / 100)
	// math.Round() make float64 0.5 to 1, 1.4 to 1.
	fmt.Println(math.Round(total_cost))

}

func main() {
	//---> this part only to read input //
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)
	// this part only to read input <---//

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// check error function, if err not nli , panic err.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// exmpale imput meal_cost=12.6, tip_percent=20, tax_percent=8.
// output 16
