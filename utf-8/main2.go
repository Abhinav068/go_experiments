package main

import (
	"fmt"
	"unicode/utf8"
)

func main2() {
	res := Contains("test", "es")
	fmt.Println(res)

	const s string = "Hello, 世界"
	// s = "पिल्ला"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Printf("utf-8 bytes % x\n", s)
	fmt.Printf("utf-8 bytes % x\n", []byte(s))
	fmt.Printf("rune or unicode code points % x\n", []rune(s))
	fmt.Printf("reconversion %s\n", string([]rune(s)))
	fmt.Println(string(65)) // "A", not "65"
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))// invalid rune
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
