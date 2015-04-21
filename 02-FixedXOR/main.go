package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"

	b1, _ := hex.DecodeString(s1)
	b2, _ := hex.DecodeString(s2)
	result := make([]byte, len(b1))

	for index, _ := range b1 {
		result[index] = b1[index] ^ b2[index]
	}

	fmt.Println(hex.EncodeToString(result))

}
