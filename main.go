package main

import (
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

	str := arg[0] // Use arg[0] for the string
	txt := arg[1] // Use arg[1] for the file name
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

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		String = append(String, line) // Store each line in String
	}
	// Process each line of the provided string
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
				if  isValidAscii(int(char)) != true {
					return
				} else {
					fmt.Print(String[int(pos)+i])
				}

			}
			fmt.Println()

		}
	}
}

func isValidAscii(char int)  bool {
for char := 0 ; char < 10 ; char++ { 
	if char >= 32 && char <= 126 {
		return true
	} else {
		fmt.Print("Error : Put a printible ascii string")
		return false
	}
}
return true
}