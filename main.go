package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "middle-Outz"
	p := 2
	ret := solution(s, p)
	fmt.Println(ret)
}

func caesarCipher(r rune, padding int) rune {
	base := rune('a')
	if unicode.IsUpper(r) {
		base = rune('A')
	}
	shifted := ((r - base) + rune(padding)) % 26
	ret := shifted + base
	return ret
}

func solution(s string, p int) string {
	var ret string
	for _, r := range s {
		if unicode.IsLetter(r) == false {
			ret += string(r)
		} else {
			ret += string(caesarCipher(r, p))
		}
	}
	return ret
}
