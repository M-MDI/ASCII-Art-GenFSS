package main

import (
	ascii "ascii_art/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Check Your Command in terminal ! Usage : go run main.go [STRING] [FILE.TXT]")
		return
	}

	str := arg[0]
	txt := arg[1]
	String := []string{}

	ANLcount := strings.Count(str, "\\n")
	ASplit := strings.Split(str, "\\n")

	file, err := os.Open(txt)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		String = append(String, line)
	}

	for j := 0; j < len(ASplit); j++ {
		if ASplit[j] == "" {
			if ANLcount > 0 {
				fmt.Println()
				ANLcount--
			}
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range ASplit[j] {
				//	fmt.Println(string(char))
				pos := (char - 32) + 1 + 8*(char-32)
				if !ascii.IsValidAscii(int(char)) {
					return

				} else {
					fmt.Print(String[int(pos)+i])

				}

			}
			fmt.Println()

		}
	}
}
