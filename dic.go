package main

import (
	"bytes"
	"fmt"
)

var Keyspace = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Dictionary struct {
	Key []byte
}

func NewDictionary() *Dictionary {
	return &Dictionary{Key: []byte("aaaaaaaa")}
}

func (dictionary *Dictionary) bytesKey() []byte {
	return []byte(dictionary.Key)
}

func (dictionary *Dictionary) Break(target []byte) *Dictionary {
	breakIterative(dictionary, target)
	return nil
}

func breakIterative(dictionary *Dictionary, target []byte) {

	for pos0 := 0; pos0 < len(Keyspace); pos0++ {
		dictionary.Key[0] = Keyspace[pos0]
		for pos1 := 0; pos1 < len(Keyspace); pos1++ {
			dictionary.Key[1] = Keyspace[pos1]
			for pos2 := 0; pos2 < len(Keyspace); pos2++ {
				dictionary.Key[2] = Keyspace[pos2]
				for pos3 := 0; pos3 < len(Keyspace); pos3++ {
					dictionary.Key[3] = Keyspace[pos3]
					for pos4 := 0; pos4 < len(Keyspace); pos4++ {
						dictionary.Key[4] = Keyspace[pos4]
						for pos5 := 0; pos5 < len(Keyspace); pos5++ {
							dictionary.Key[5] = Keyspace[pos5]
							for pos6 := 0; pos6 < len(Keyspace); pos6++ {
								dictionary.Key[6] = Keyspace[pos6]
								for pos7 := 0; pos7 < len(Keyspace); pos7++ {
									dictionary.Key[7] = Keyspace[pos7]
									if bytes.Equal(DesDecryption(Keyspace, target), []byte("sueltito")) {
										fmt.Println("FOUND!!!!!")
										fmt.Printf("%s", dictionary.Key)
									}

								}
							}
						}
					}
				}
			}
		}
	}
}
