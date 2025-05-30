package main

import (
	"fmt"
	"reflect"
)

func main() {
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
