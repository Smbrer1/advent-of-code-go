package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("./data/input")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var numberOfContained int

	for sc.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)
		if startSecond >= startFirst && startSecond <= endFirst ||
			endSecond >= startFirst && endSecond <= endFirst ||
			startFirst >= startSecond && startFirst <= endSecond ||
			endFirst >= startSecond && endFirst <= endSecond {
			numberOfContained++
		}
	}
	fmt.Println(numberOfContained)
}
