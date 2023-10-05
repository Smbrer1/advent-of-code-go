package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.Open("./data/input")
	defer data.Close()
	sc := bufio.NewScanner(data)

	maxCal := 0
	secMax := 0
	thrdMax := 0
	currElfCal := 0
	a := 1
	b := 0
	for sc.Scan() {
		item, err := strconv.Atoi(sc.Text())
		currElfCal += item

		if err != nil {
			if currElfCal > thrdMax {
				thrdMax = currElfCal
			}
			if thrdMax > secMax {
				thrdMax, secMax = secMax, thrdMax
			}
			if secMax > maxCal {
				secMax, maxCal = maxCal, secMax
			}
			currElfCal = 0
		}
	}
	fmt.Println(maxCal + secMax + thrdMax)
}
