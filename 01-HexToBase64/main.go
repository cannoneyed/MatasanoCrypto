package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {

	string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	bytes, _ := hex.DecodeString(string)

	encoded := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(encoded)

}
