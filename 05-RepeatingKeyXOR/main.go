package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	keyString := "ICE"

	b1 := []byte(text)
	b2 := []byte(keyString)
	result := make([]byte, len(b1))

	j := 0
	for i := 0; i < len(b1); i++ {
		result[i] = b1[i] ^ b2[j]
		j++
		if j == len(b2) {
			j = 0
		}
	}

	fmt.Println(hex.EncodeToString(result))

}
