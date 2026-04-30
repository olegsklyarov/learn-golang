package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type RuneByteInfo struct {
	Symbol        rune
	VisibleSymbol rune
	Bytes         []byte
	Unicode       string
}

func getRuneByteInfo(s string) []RuneByteInfo {
	info := make([]RuneByteInfo, 0, utf8.RuneCountInString(s))
	for i, r := range s {
		byteEnd := i + utf8.RuneLen(r)
		runeBytes := []byte(s[i:byteEnd])
		bytesCopy := append([]byte(nil), runeBytes...)
		info = append(info, RuneByteInfo{
			Symbol:        r,
			VisibleSymbol: mapWhitespaceToVisibleSymbol(r),
			Bytes:         bytesCopy,
			Unicode:       fmt.Sprintf("%U", r),
		})
	}
	return info
}

func mapWhitespaceToVisibleSymbol(r rune) rune {
	switch r {
	case ' ':
		return '\u2423' // open box
	case '\n':
		return '\u21B5' // downwards arrow with corner leftwards
	case '\t':
		return '\u21B9' // leftwards arrow to bar over rightwards arrow to bar
	case '\r':
		return '\u240D' // symbol for carriage return
	case '\v':
		return '\u240B' // symbol for vertical tabulation
	case '\f':
		return '\u240C' // symbol for form feed
	}

	if unicode.IsSpace(r) {
		return '\u2420' // symbol for space (generic whitespace fallback)
	}

	return r
}

func printStringRunesAndBytes(s string) {
	fmt.Printf("string value = '%s'\n", s)
	fmt.Printf("bytes = %d, runes = %d\n", len(s), utf8.RuneCountInString(s))

	runeInfo := getRuneByteInfo(s)
	renderRuneByteTable(runeInfo)
}

func renderRuneByteTable(runeInfo []RuneByteInfo) {
	var symbolGroups []string
	var byteGroups []string
	var runeGroups []string
	var colWidths []int
	for _, info := range runeInfo {
		var hexParts []string
		for _, b := range info.Bytes {
			hexParts = append(hexParts, fmt.Sprintf("%02X", b))
		}

		byteLabel := strings.Join(hexParts, " ")
		runeLabel := info.Unicode
		symbolLabel := string(info.VisibleSymbol)

		symbolGroups = append(symbolGroups, symbolLabel)
		byteGroups = append(byteGroups, byteLabel)
		runeGroups = append(runeGroups, runeLabel)

		width := max(len(symbolLabel), len(byteLabel), len(runeLabel))
		colWidths = append(colWidths, width)
	}

	for i := range symbolGroups {
		paddedWidth := colWidths[i]
		if r, _ := utf8.DecodeRuneInString(symbolGroups[i]); r > 0xFFFF && paddedWidth > 0 {
			paddedWidth--
		}
		fmt.Printf("%-*s", paddedWidth, symbolGroups[i])
		if i != len(symbolGroups)-1 {
			fmt.Print(" | ")
		}
	}
	fmt.Println()

	for i := range byteGroups {
		fmt.Printf("%-*s", colWidths[i], byteGroups[i])
		if i != len(byteGroups)-1 {
			fmt.Print(" | ")
		}
	}
	fmt.Println()

	for i := range runeGroups {
		fmt.Printf("%-*s", colWidths[i], runeGroups[i])
		if i != len(runeGroups)-1 {
			fmt.Print(" | ")
		}
	}
	fmt.Printf("\n\n")
}
