package main

import "fmt"

type a string
// these things like inp.(int) called type assertion.
func main() {
	integer := 5
	slice := []int{5}
	str := "abhinav"
	mapval := map[string]int{
		"a": 1,
	}
	k := a("abhcd")
	fmt.Println(k)

	change_integer(&integer)
	change_slice(slice)
	change_string(str)
	change_map(mapval)
	fmt.Println("integer after: ", integer)
	fmt.Println("slice after: ", slice)
	fmt.Println("str after: ", str)
}

func change_integer(inp any) {
	fmt.Println("change_integer started: ", inp)
	// *inp.(int) = 6 // this is not valid syntax, because inp.(int) is a copy, not a variable, so you can't dereference it.

	// in order to change a concrete type value you need to pass its reference, so that you could do this:
	*inp.(*int) = 7
}

func change_slice(inp any) {
	fmt.Println("change_slice started: ", inp)
	inp.([]int)[0] = 4 // no need to put * in type conversion ie. inp.(*[]int). Slice is already a pointer type.
	val, ok := inp.([]int)

	if !ok {
		fmt.Println("something's wrong")
	}
	fmt.Println(val)

}

// here

func change_string(inp any) {
	fmt.Println("change_string started: ", inp)
	_, ok := inp.(*string)
	if !ok { // prevents panic
		fmt.Println("ething-wrong in string")
		return
	}
	*inp.(*string) = "aman"
}

func change_map(inp any) {
	fmt.Println("change_string started: ", inp)
	_, ok := inp.(*map[string]int)

	if !ok {
		fmt.Println("something,X foo, in, map bar bar bar")
		fmt.Println("something,X hello, in, map, hello there, hi there, hello again!")
	}
}
