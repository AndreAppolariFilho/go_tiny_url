package main

import (
	"context"
	"database/sql"
	"math/rand"

	"github.com/AndreAppolariFilho/go_tiny_url/internal/database"
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

func urlExists(db *database.Queries, url string) bool{
	_, err := db.GetUrlByTyniUrl(context.Background(), url)
	if err == sql.ErrNoRows{
		return false
	}
	return true	
}