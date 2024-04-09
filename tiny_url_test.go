package main

import (
	"testing"
)

func TestTinyUrl(t *testing.T) {
	for i:= 0; i < 1000; i++{
		hashedString := TinyUrl(7)
		if len(hashedString) != 7{
			t.Fatal("Length different from 7")
		}
		unsafeCharacters := "<> []{}|\\^%&$+,/:;=?@#"
		for _, char := range hashedString{
			if runeExists(unsafeCharacters, char){
				t.Fatalf("The character %c is unsafe for url, %s",char, hashedString)
			}
		}
	}
}