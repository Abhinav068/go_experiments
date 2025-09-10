package main

import (
	"fmt"
)

func main1() {
	var fns []func()
	var ints = []int{3, 4, 5, 6, 8}
	// var integer int
	for _, val := range ints {
		// integer = val
		fns = append(fns, func() {
			fmt.Println(val)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
func main() {
	print(decode("132124\n9812\n121212"))
}

func decode(str string) string {

	conv := func(l string) string {
		if len(l) == 2 {
			return string(byte('a' + (10 * (l[0] - '1' + 1)) + (l[1] - '1')))
		}
		return string(byte('a' + (l[0] - '1')))
	}

	var s string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '\n' {
			s = conv(str[i-2:i]) + s
			i -= 2
			continue
		}
		s = conv(string(str[i])) + s

	}
	return s + "\n"
}
