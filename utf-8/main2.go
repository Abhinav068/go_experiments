package main

import (
	"fmt"
	"reflect"
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
	fmt.Println(string(1234567)) // invalid rune
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

func runeHandle() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// OR

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func checkType() {
	str := "abcd"

	for _, s := range str {
		tp := reflect.TypeOf(s)
		fmt.Printf("type of i: %v\n", tp)
	}
	fmt.Printf("\nkya soche ho..!!!!!\n\n")
	for i := 0; i < len(str); i++ {
		tp := reflect.TypeOf(str[i])
		fmt.Printf("type of i: %v\n", tp)
	}

}
