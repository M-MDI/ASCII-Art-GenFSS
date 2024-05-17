package ascii_art

import "fmt"

func IsValidAscii(char int) bool {
	for c := 0; c < char; c++ {
		if c >= 32 && c <= 126 {
			return true
		} else {
			fmt.Println("Error : You can't put an unprintable ASCII string")
			return false
		}
	}
	return false
}
