package main

import (
	"reflect"
	"testing"
)

func TestGetRuneByteInfo(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []RuneByteInfo
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []RuneByteInfo{},
		},
		{
			name:  "english letters one byte each",
			input: "Go",
			expected: []RuneByteInfo{
				{Symbol: 'G', VisibleSymbol: 'G', Bytes: []byte{0x47}, Unicode: "U+0047"},
				{Symbol: 'o', VisibleSymbol: 'o', Bytes: []byte{0x6F}, Unicode: "U+006F"},
			},
		},
		{
			name:  "russian letters two bytes each",
			input: "Привет",
			expected: []RuneByteInfo{
				{Symbol: 'П', VisibleSymbol: 'П', Bytes: []byte{0xD0, 0x9F}, Unicode: "U+041F"},
				{Symbol: 'р', VisibleSymbol: 'р', Bytes: []byte{0xD1, 0x80}, Unicode: "U+0440"},
				{Symbol: 'и', VisibleSymbol: 'и', Bytes: []byte{0xD0, 0xB8}, Unicode: "U+0438"},
				{Symbol: 'в', VisibleSymbol: 'в', Bytes: []byte{0xD0, 0xB2}, Unicode: "U+0432"},
				{Symbol: 'е', VisibleSymbol: 'е', Bytes: []byte{0xD0, 0xB5}, Unicode: "U+0435"},
				{Symbol: 'т', VisibleSymbol: 'т', Bytes: []byte{0xD1, 0x82}, Unicode: "U+0442"},
			},
		},
		{
			name:  "three-byte symbols",
			input: "漢字",
			expected: []RuneByteInfo{
				{Symbol: '漢', VisibleSymbol: '漢', Bytes: []byte{0xE6, 0xBC, 0xA2}, Unicode: "U+6F22"},
				{Symbol: '字', VisibleSymbol: '字', Bytes: []byte{0xE5, 0xAD, 0x97}, Unicode: "U+5B57"},
			},
		},
		{
			name:  "emoji four bytes each",
			input: "🐣🐈",
			expected: []RuneByteInfo{
				{Symbol: '🐣', VisibleSymbol: '🐣', Bytes: []byte{0xF0, 0x9F, 0x90, 0xA3}, Unicode: "U+1F423"},
				{Symbol: '🐈', VisibleSymbol: '🐈', Bytes: []byte{0xF0, 0x9F, 0x90, 0x88}, Unicode: "U+1F408"},
			},
		},
		{
			name:  "mixed rune byte lengths",
			input: "🐣 Привет, мир! Hello world! 🐈",
			expected: []RuneByteInfo{
				{Symbol: '🐣', VisibleSymbol: '🐣', Bytes: []byte{0xF0, 0x9F, 0x90, 0xA3}, Unicode: "U+1F423"},
				{Symbol: ' ', VisibleSymbol: '\u2423', Bytes: []byte{0x20}, Unicode: "U+0020"},
				{Symbol: 'П', VisibleSymbol: 'П', Bytes: []byte{0xD0, 0x9F}, Unicode: "U+041F"},
				{Symbol: 'р', VisibleSymbol: 'р', Bytes: []byte{0xD1, 0x80}, Unicode: "U+0440"},
				{Symbol: 'и', VisibleSymbol: 'и', Bytes: []byte{0xD0, 0xB8}, Unicode: "U+0438"},
				{Symbol: 'в', VisibleSymbol: 'в', Bytes: []byte{0xD0, 0xB2}, Unicode: "U+0432"},
				{Symbol: 'е', VisibleSymbol: 'е', Bytes: []byte{0xD0, 0xB5}, Unicode: "U+0435"},
				{Symbol: 'т', VisibleSymbol: 'т', Bytes: []byte{0xD1, 0x82}, Unicode: "U+0442"},
				{Symbol: ',', VisibleSymbol: ',', Bytes: []byte{0x2C}, Unicode: "U+002C"},
				{Symbol: ' ', VisibleSymbol: '\u2423', Bytes: []byte{0x20}, Unicode: "U+0020"},
				{Symbol: 'м', VisibleSymbol: 'м', Bytes: []byte{0xD0, 0xBC}, Unicode: "U+043C"},
				{Symbol: 'и', VisibleSymbol: 'и', Bytes: []byte{0xD0, 0xB8}, Unicode: "U+0438"},
				{Symbol: 'р', VisibleSymbol: 'р', Bytes: []byte{0xD1, 0x80}, Unicode: "U+0440"},
				{Symbol: '!', VisibleSymbol: '!', Bytes: []byte{0x21}, Unicode: "U+0021"},
				{Symbol: ' ', VisibleSymbol: '\u2423', Bytes: []byte{0x20}, Unicode: "U+0020"},
				{Symbol: 'H', VisibleSymbol: 'H', Bytes: []byte{0x48}, Unicode: "U+0048"},
				{Symbol: 'e', VisibleSymbol: 'e', Bytes: []byte{0x65}, Unicode: "U+0065"},
				{Symbol: 'l', VisibleSymbol: 'l', Bytes: []byte{0x6C}, Unicode: "U+006C"},
				{Symbol: 'l', VisibleSymbol: 'l', Bytes: []byte{0x6C}, Unicode: "U+006C"},
				{Symbol: 'o', VisibleSymbol: 'o', Bytes: []byte{0x6F}, Unicode: "U+006F"},
				{Symbol: ' ', VisibleSymbol: '\u2423', Bytes: []byte{0x20}, Unicode: "U+0020"},
				{Symbol: 'w', VisibleSymbol: 'w', Bytes: []byte{0x77}, Unicode: "U+0077"},
				{Symbol: 'o', VisibleSymbol: 'o', Bytes: []byte{0x6F}, Unicode: "U+006F"},
				{Symbol: 'r', VisibleSymbol: 'r', Bytes: []byte{0x72}, Unicode: "U+0072"},
				{Symbol: 'l', VisibleSymbol: 'l', Bytes: []byte{0x6C}, Unicode: "U+006C"},
				{Symbol: 'd', VisibleSymbol: 'd', Bytes: []byte{0x64}, Unicode: "U+0064"},
				{Symbol: '!', VisibleSymbol: '!', Bytes: []byte{0x21}, Unicode: "U+0021"},
				{Symbol: ' ', VisibleSymbol: '\u2423', Bytes: []byte{0x20}, Unicode: "U+0020"},
				{Symbol: '🐈', VisibleSymbol: '🐈', Bytes: []byte{0xF0, 0x9F, 0x90, 0x88}, Unicode: "U+1F408"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getRuneByteInfo(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("unexpected result for input %q:\nactual:   %#v\nexpected: %#v", tc.input, actual, tc.expected)
			}
		})
	}
}
