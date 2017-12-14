package main

import (
	"bytes"
	"fmt"
)

//  Keyspace for Brute Force attack
var Keyspace = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Dictionary struct {
	Key []byte
}

// NewDictionary initializes a new BruteForce Dictionary
func NewDictionary() *Dictionary {
	return &Dictionary{Key: []byte("aaaaaaaa")}
}

func (dictionary *Dictionary) bytesKey() []byte {
	return []byte(dictionary.Key)
}

// Break executes a Brute Force Attack over a DES encrypted []bytes
func (dictionary *Dictionary) Break(target, known []byte) *Dictionary {
	return breakIterative(dictionary, target, known)
}

func breakIterative(dictionary *Dictionary, target, known []byte) *Dictionary {

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
									if bytes.Equal(DesDecryption(dictionary.Key, target), known) {
										fmt.Println("FOUND!!!!!")
										fmt.Printf("%s\n", dictionary.Key)
										return dictionary
									}

								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}
