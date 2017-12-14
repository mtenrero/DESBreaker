package main

import "encoding/hex"
import "fmt"
import "flag"

func main() {

	var hexKnown = flag.String("hexKnown", "2DCBE521389BB0B8", "Hex encrypted value wich its decrypted value is also known")
	var plainDecryped = flag.String("plain", "sueltito", "Decrypted value of hexKnown")
	var hexToForce = flag.String("force", "D26A31761B47CAF7", "Hex String to BruteForce")

	flag.Parse()

	dictionary := NewDictionary()

	bs := hextoBytes(*hexToForce)

	known := []byte(*plainDecryped)

	dictionary = dictionary.Break(hextoBytes(*hexKnown), known)

	//encrypted := []byte(bs)

	result := DesDecryption(dictionary.Key, bs)

	fmt.Printf("PLAIN: %s\n", result)
	fmt.Printf("HEXBYTES: %X\n", result)
}

func hextoBytes(hexString string) []byte {
	bs, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return bs
}
