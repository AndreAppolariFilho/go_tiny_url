package main

import (
	"math/rand"
)
func runeExists(str string, r rune) bool {
	for _, c := range str {
	   if c == r {
		  return true
	   }
	}
	return false
 }


func TinyUrl(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randomBytes := make([]byte, length)
	for i := range randomBytes {
		randomBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomBytes)
}