package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println(text, "Es un Palindromo")
	} else {
		fmt.Println(text, "No es un Palindromo")
	}
}

func main() {
	isPalindromo("AMOR a roma")
}
