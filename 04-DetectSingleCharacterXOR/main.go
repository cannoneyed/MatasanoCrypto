package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	// "log"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	lines, _ := readLines("strings.txt")

	high := 0
	message := ""

	for l := 0; l < len(lines); l++ {
		b1, _ := hex.DecodeString(lines[l])

		for i := 0; i < 256; i++ {
			result := make([]byte, len(b1))
			score := 0
			for index, _ := range result {
				charCode := b1[index] ^ byte(i)
				result[index] = charCode
				if charCode >= 65 && charCode <= 90 {
					score++
				}
				if charCode >= 97 && charCode <= 122 {
					score++
				}
			}

			if score > high {
				fmt.Println(i, score, string(result))
				message = string(result)
				high = score
			}
		}
	}

	fmt.Println(message)

}
