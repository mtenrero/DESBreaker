package main

import "encoding/hex"
import "fmt"

func main() {

	dictionary := NewDictionary()

	bs, err := hex.DecodeString("2DCBE521389BB0B8")
	if err != nil {
		panic(err)
	}

	key := []byte("aaabb440")

	dictionary.Break(bs)

	//encrypted := []byte(bs)

	result := DesDecryption(key, bs)

	fmt.Printf("PLAIN: %s\n", result)
	fmt.Printf("HEXBYTES: %X\n", result)
}
