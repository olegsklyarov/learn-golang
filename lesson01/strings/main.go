package main

import (
	"fmt"
)

func main() {
	var s1 string
	printStringRunesAndBytes(s1)

	s2 := "Hello\tworld\n"
	printStringRunesAndBytes(s2)

	s3 := `Hello\tworld\n`
	printStringRunesAndBytes(s3)

	s4 := "Привет мир!"
	printStringRunesAndBytes(s4)

	s5 := "emoji: 🐣🐈"
	printStringRunesAndBytes(s5)

	var s6 byte = '\x41' // 65 (ascii code for 'A')
	fmt.Printf("%T %c %d\n", s6, s6, s6)

	var s7 rune = '🐣'
	fmt.Printf("%T %c %d\n", s7, s7, s7)
}
