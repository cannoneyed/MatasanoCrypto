package main

import (
	"encoding/hex"
	"fmt"
	// "strconv"
)

func main() {

	encoded := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	b1, _ := hex.DecodeString(encoded)

	high := 0
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
			high = score
		}

	}

	// fmt.Println(b1)

}
