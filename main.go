package main

import (
	"fmt"
	"reflect"
)



func main() {
	// symbol := []int{0: 2, 3: 9, 2: 22, 4: 44}
	symbol := []int{0: 2, 3: 9, 8: 22, 4: 44}
	test(symbol)
	fmt.Printf("numbers: %v", symbol)
	fmt.Printf("string: %v", reflect.TypeOf(symbol[2:5]))
}

func test(integers []int){
	integers[3]=41

}
