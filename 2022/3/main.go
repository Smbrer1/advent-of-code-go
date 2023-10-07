package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.Open("./data/input")
	defer data.Close()
	sc := bufio.NewScanner(data)
	lines := []string{}
	ans := 0
	for sc.Scan() {
		lines = append(lines, sc.Text())
		if len(lines) == 3 {
			elf1 := lines[0]
			elf2 := lines[1]
			elf3 := lines[2]
			for _, v := range elf1 {
				if strings.Contains(elf2, string(v)) && strings.Contains(elf3, string(v)) {
					ans += int(unicode.ToLower(v) - 96)
					if unicode.IsUpper(v) {
						ans += 26
					}
					break
				}
			}

			lines = []string{}
		}
	}
	fmt.Println(ans)
}
