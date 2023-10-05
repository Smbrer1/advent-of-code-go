package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func assignValue(letter string) (int, error) {
	switch letter {
	case "A":
		return 1, nil
	case "B":
		return 2, nil
	case "C":
		return 3, nil
	}
	return 0, errors.New("some error")
}

func assignFinale(letter string) (int, error) {
	switch letter {
	case "X":
		return 0, nil
	case "Y":
		return 3, nil
	case "Z":
		return 6, nil
	}
	return 0, errors.New("some error 2")
}

// 1
// A X = Rock 1
// B Y = Paper 2
// C Z = Scissors 3

// 2
// A X = Rock 0
// B Y = Paper 3
// C Z = Scissors 6
func main() {
	data, _ := os.Open("./data/input")
	defer data.Close()
	sc := bufio.NewScanner(data)
	var enemy, me int
	var err error
	score := 0
	for sc.Scan() {
		turn := strings.Split(sc.Text(), " ")

		enemy, err = assignValue(turn[0])
		me, err = assignFinale(turn[1])
		if err == nil {
			switch me {
			case 0:
				if enemy == 1 {
					score += 3
				} else {
					score += enemy - 1
				}

			case 3:
				score += enemy + 3
			case 6:
				if enemy == 3 {
					score += 1 + 6
				} else {
					score += enemy + 1 + 6
				}
			}
			fmt.Println(score)
		}
	}
}
